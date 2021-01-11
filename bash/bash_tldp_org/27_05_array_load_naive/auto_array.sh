#!/bin/bash
# script-array.sh: Loads this script into an array.
# Inspired by an e-mail from Chris Martin (thanks!).
declare -a script
idx=0

# `mapfile' can be used in this case, instead of this:
while IFS= read -r line;do
    script[idx]="$line"
    ((idx++))
done < "$0"

echo

for (( i=0; i<${#script[@]}; i++ )) { echo "${script[i]}"; }

exit $?

# Exercise:
# --------
#  Modify this script so it lists itself
#+ in its original format,
#+ complete with whitespace, line breaks, etc.
