#!/bin/bash
# de-rpm.sh: Unpack an 'rpm' archive
USAGE="Usage: $(basename $0) <target-rpm-archive>"

# tests
${1?"$USAGE"}
${2?"$USAGE"}
[[ ! -e "$2" || ${2:(-4)} != ".rpm") ]] && echo "$USAGE"


TEMPFILE=$$.cpio                         #  Tempfile with "unique" name.
                                         #  $$ is process ID of script.

rpm2cpio < $1 > $TEMPFILE                #  Converts rpm archive into
                                         #+ cpio archive.
cpio --make-directories -F $TEMPFILE -i  #  Unpacks cpio archive.
rm -f $TEMPFILE                          #  Deletes cpio archive.

exit $?

#  Exercise:
#  Add check for whether 1) "target-file" exists and
#+                       2) it is an rpm archive.
#  Hint:                    Parse output of 'file' command.
