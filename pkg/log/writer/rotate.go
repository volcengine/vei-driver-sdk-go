/*
 * Copyright 2023 Beijing Volcano Engine Technology Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This file is modified based on https://github.com/influxdata/telegraf/blob/v1.27.0/internal/rotate/file_writer.go

package writer

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	FilePerm   = os.FileMode(0644)
	DateFormat = "2006-01-02"
)

type FileWriter struct {
	filename                 string
	filenameRotationTemplate string
	current                  *os.File
	interval                 time.Duration
	maxArchives              int
	expireTime               time.Time
	sync.Mutex
}

// NewFileWriter creates a new file writer.
func NewFileWriter(filename string, maxArchives int) (io.WriteCloser, error) {
	w := &FileWriter{
		filename:                 filename,
		interval:                 time.Hour * 24,
		maxArchives:              maxArchives,
		filenameRotationTemplate: getFilenameRotationTemplate(filename),
	}

	if err := w.openCurrent(); err != nil {
		return nil, err
	}

	return w, nil
}

func openFile(filename string) (*os.File, error) {
	if err := os.MkdirAll(path.Dir(filename), 0775); err != nil {
		return nil, err
	}
	return os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, FilePerm)
}

func getFilenameRotationTemplate(filename string) string {
	// Extract the file extension
	fileExt := filepath.Ext(filename)
	// Remove the file extension from the filename (if any)
	stem := strings.TrimSuffix(filename, fileExt)
	return stem + ".%s" + fileExt
}

// Write writes p to the current file, then checks to see if
// rotation is necessary.
func (w *FileWriter) Write(p []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	if err = w.rotateIfNeeded(); err != nil {
		return 0, err
	}

	if n, err = w.current.Write(p); err != nil {
		return 0, err
	}

	return n, nil
}

// Close closes the current file.  Writer is unusable after this
// is called.
func (w *FileWriter) Close() (err error) {
	w.Lock()
	defer w.Unlock()

	// Rotate before closing
	if err := w.rotateIfNeeded(); err != nil {
		return err
	}

	// Close the file if we did not rotate
	if err := w.current.Close(); err != nil {
		return err
	}

	w.current = nil
	return nil
}

func (w *FileWriter) openCurrent() (err error) {
	year, month, day := time.Now().Date()
	w.expireTime = time.Date(year, month, day, 0, 0, 0, 0, time.Local).Add(w.interval)
	w.current, err = openFile(w.filename)
	if err != nil {
		return err
	}
	return w.rotateIfNeeded()
}

func (w *FileWriter) rotateIfNeeded() error {
	if w.interval > 0 && time.Now().After(w.expireTime) {
		if err := w.rotate(); err != nil {
			//Ignore rotation errors and keep the log open
			fmt.Printf("unable to rotate the file %q, %s", w.filename, err.Error())
		}
		return w.openCurrent()
	}
	return nil
}

func (w *FileWriter) rotate() (err error) {
	if err = w.current.Close(); err != nil {
		return err
	}

	// Use year-month-date for readability, unix time to make the file name unique with second precision
	now := time.Now()
	rotatedFilename := fmt.Sprintf(w.filenameRotationTemplate, now.Format(DateFormat))
	if err = os.Rename(w.filename, rotatedFilename); err != nil {
		return err
	}

	return w.purgeArchivesIfNeeded()
}

func (w *FileWriter) purgeArchivesIfNeeded() (err error) {
	if w.maxArchives == -1 {
		//Skip archiving
		return nil
	}

	var matches []string
	if matches, err = filepath.Glob(fmt.Sprintf(w.filenameRotationTemplate, "*")); err != nil {
		return err
	}

	//if there are more archives than the configured maximum, then purge older files
	if len(matches) > w.maxArchives {
		//sort files alphanumerically to delete older files first
		sort.Strings(matches)
		for _, filename := range matches[:len(matches)-w.maxArchives] {
			if err = os.Remove(filename); err != nil {
				return err
			}
		}
	}
	return nil
}
