#!/bin/bash
# findstring.sh:
# Find a particular string in the binaries in a specified directory.

#directory=/usr/bin/
#fstring="Free Software Foundation"  # See which files come from the FSF.
ERR_MSG="Usage: $(basename $0) <directory-to-search> <string-to-search>"

directory=${1?"$ERR_MSG"}  # beautiful bash idioms: exit with error and message
fstring=${2?"$ERR_MSG"}

for file in $( find $directory -type f -name 'z*' | sort ); do
  strings -f $file | grep "$fstring" | sed -e "s%$1%%"
  #  In the "sed" expression,
  #+ it is necessary to substitute for the normal "/" delimiter
  #+ because "/" happens to be one of the characters filtered out.
  #  Failure to do so gives an error message. (Try it.)
done

exit $?

#  Exercise (easy):
#  ---------------
#  Convert this script to take command-line parameters
#+ for $directory and $fstring.
