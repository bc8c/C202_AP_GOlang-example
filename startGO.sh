#!/bin/bash

if [ $# -ne 1 ];then
  echo "#usage: $0 <gofile>"
  exit
fi

FILE=$(echo $1 | cut -f 1 -d '.')
go build $1
./$FILE
rm $FILE