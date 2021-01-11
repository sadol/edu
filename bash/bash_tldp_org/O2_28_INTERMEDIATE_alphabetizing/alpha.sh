#!/usr/bin/bash

# INFO: locale!!!
output=( "$(sort -i < <( for((i=0;i<${#1};i++)) { echo ${1:i:1}; } ) )" )
echo ${output[@]}
