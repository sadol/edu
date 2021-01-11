#!/bin/bash
# rn.sh

# Very simpleminded filename "rename" utility (based on "lowercase.sh").
#
#  The "ren" utility, by Vladimir Lanin (lanin@csd2.nyu.edu),
#+ does a much better job of this.

ARGS=2
E_BADARGS=85
ONE=1                          # For getting singular/plural right (see below).

if [[ $# -ne $ARGS ]] ;then
    echo "Usage: $(basename $0) old-pattern new-pattern"
    # As in "rn gif jpg", which renames all gif files in working directory to jpg.
    exit $E_BADARGS
fi

number=0                      # Keeps track of how many files actually renamed.

for filename in *$1* ;do             #Traverse all matching files in directory.
    if [[ -f "$filename" || -d "$filename" && -x "$filename" ]];then
        fname="$(basename $filename)"                         # Strip off path.
        n=$(sed -nr 's/'"$1"'/'"$2"'/p' <<< "$fname")
        mv "$fname" "$n"
        ((number++))
    fi
done

if [[ "$number" -eq "$ONE" ]];then                       # For correct grammar.
    echo "$number file renamed."
else
    echo "$number files renamed."
fi

exit $?

# Exercises:
# ---------
# What types of files will this not work on?
# How can this be fixed?
