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

package writer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileWriter_NoRotation(t *testing.T) {
	tempDir := t.TempDir()
	writer, err := NewFileWriter(filepath.Join(tempDir, "test"), 0)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, writer.Close()) })

	_, err = writer.Write([]byte("Hello World"))
	require.NoError(t, err)
	_, err = writer.Write([]byte("Hello World 2"))
	require.NoError(t, err)
	files, _ := os.ReadDir(tempDir)
	require.Equal(t, 1, len(files))
}

func TestFileWriter_CloseDoesNotRotate(t *testing.T) {
	tempDir := t.TempDir()
	writer, err := NewFileWriter(filepath.Join(tempDir, "test.log"), -1)
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	files, _ := os.ReadDir(tempDir)
	require.Equal(t, 1, len(files))
	require.Regexp(t, "^test.log$", files[0].Name())
}
