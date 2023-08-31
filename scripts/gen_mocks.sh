#! /bin/bash

DIR="."

function process() {
    for file in `ls $1`
    do
      sub_path="$1/$file"

      if [ ! -d $sub_path ];then
        continue
      fi

      if [ -d "$sub_path/interfaces" ]; then
        echo -e "\033[33mauto generate mocks for interfaces in $sub_path \033[0m"
        mockery --dir "$sub_path/interfaces" --all --output "$sub_path/interfaces/mocks"
      fi

      process $sub_path
    done
}

process $DIR
