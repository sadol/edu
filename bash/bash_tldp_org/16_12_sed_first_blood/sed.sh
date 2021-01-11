#!/bin/bash
# wf.sh: Crude word frequency analysis on a text file.
# This is a more efficient version of the "wf2.sh" script.


# Check for input file on command-line.
ARGS=1
E_BADARGS=85
E_NOFILE=86

if [[ $# -ne "$ARGS" ]];then  # Correct number of arguments passed to script?
    echo "Usage: $(basename $0) filename"
    exit $E_BADARGS
fi

if [[ ! -f "$1" ]];then       # Check if file exists.
    echo "File \"$1\" does not exist."
    exit $E_NOFILE
fi


########################################################
# main ()
#sed -e 's/\.//g'  -e 's/\,//g' -e 's/ /\
#  Arun Giridhar suggests modifying the above to:
#  . . . | sort | uniq -c | sort +1 [-f] | sort +0 -nr
#  This adds a secondary sort key, so instances of
#+ equal occurrence are sorted alphabetically.
#  As he explains it:
#  "This is effectively a radix sort, first on the
#+ least significant column
#+ (word or string, optionally case-insensitive)
#+ and last on the most significant column (frequency)."
#
#  As Frank Wang explains, the above is equivalent to
#+       . . . | sort | uniq -c | sort +0 -nr
#+ and the following also works:
#+       . . . | sort | uniq -c | sort -k1nr -k


#sed -e 's/[[:punct:]]//g' -e 's/[[:blank:]]//g' "$1" \
#    | tr '[[:upper:]]' '[[:lower:]]' | sort | uniq -c | sort -nr
text="$(<$1)"                                  # load text file into a variable
sed -nr 's/[[:punct:][:blank:]]//gp' <<< "${text,,}" | sort | uniq -c | sort -nr
# lower-case the whole thing using var subst ^^^^^
########################################################

exit 0

# Exercises:
# ---------
# 1) Add 'sed' commands to filter out other punctuation,
#+   such as semicolons.
# 2) Modify the script to also filter out multiple spaces and
#+   other whitespace.
