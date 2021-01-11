#!/bin/bash
# script-detector.sh: Detects scripts within a directory.

ISSCRIPTCHARS=2    # Test first 2 characters.
ISBASH="bash"
SHABANG='#!'   # Scripts begin with a "sha-bang."
E_NOOFARGS=95
E_NODIR=96
NOOFARGS=1

if [[ $# -ne $NOOFARGS ]];then
    echo "Wrong number of arguments."
    echo "Usage: $(basename $0) <dir-to-search> ."
    exit $E_NOOFARGS
fi

if [[ ! -d $1 ]];then
    echo "<$1> is not a directory."
    echo "Usage: $(basename $0) <dir-to-search> ."
    exit $E_NODIR
fi

DIRTOSEARCH="$1*"

for file in $DIRTOSEARCH;do
    if [[ -f "$file" ]];then
        if [[ $(head -c$ISSCRIPTCHARS "$file") = "$SHABANG" ]];then
            if [[ $(head -1 "$file" | grep $ISBASH) ]];then
                echo "File \"$file\" is a BASH script."
            else
                echo "File \"$file\" is an unknown script."
            fi
        else
            echo "File \"$file\" is *not* a script."
        fi
    fi
done

exit 0

#  Exercises:
#  ---------
#  1) Modify this script to take as an optional argument
#+    the directory to scan for scripts
#+    (rather than just the current working directory).
#
#  2) As it stands, this script gives "false positives" for
#+    Perl, awk, and other scripting language scripts.
#     Correct this.
