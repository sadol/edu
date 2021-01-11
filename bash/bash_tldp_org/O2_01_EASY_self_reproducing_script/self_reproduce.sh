#!/usr/bin/bash

#cp $0 ${0%.*}.backup
cp $0 ${0/%sh/backup.sh}                                   # suffix replacement
exit $?
