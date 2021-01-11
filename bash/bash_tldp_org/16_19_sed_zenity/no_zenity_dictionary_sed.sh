#!/bin/bash
# dict-lookup.sh

E_BADARGS=85
E_NOTERM=86
E_NODICT=87
E_NODICTDIR=88
MAXCONTEXTLINES=50                        # Maximum number of lines to show.
DICTDIR="dicts"
DEFAULT_DICTFILE="${DICTDIR}/webster.txt"
                                          # Default dictionary file pathname.
                                          # Change this as necessary.
DEFAULT_SEDSCRIPT=""
TERMTOSEARCH=""
USAGEMSG="Usage: $(basename ${0}) Word-to-search [dictionary-file]
       Examples: AB-, aB-, Ab-, ab- etc."
DICTFILE=""
SEDSCRIPT=""

# check dictdir first
if [[ ! -d ${DICTDIR} ]];then
    echo "Dictionaries directory <${DICTDIR}> is not present.
          Directory should contain subdirs with text based dictionary and sed
          script to operate on said dictionary."
    exit $E_NODICTDR
fi

if [[ $# -lt 1 ||  $# -gt 2 ]];then
    echo "$USAGEMSG"
    exit $E_BADARGS
else
    if [[ -z "$2" ]]; then
        DICTFILE="$DEFAULT_DICTFILE"
    else
        DICTFILE="$2"
    fi
fi

if [[ ! -f $DICTFILE ]];then
    echo "There is no dictionary <${DICTFILE}>."
    exit $E_NODICT
fi

TERMTOSEARCH="$(tr [[:lower:]] [[:upper:]] <<< $1)"
if [[ -z $(sed -n '/[[:upper:][:punct:]]/ p' <<< ${TERMTOSEARCH}) ]];then
    echo "There is no term <${TERMTOSEARCH}> in dictionary <${DICTFILE}>."
    exit $E_NOTERM
fi

# webster sed script with argument
DEFAULT_SEDSCRIPT='
    # search only in the first section of the fgrep results
    /^'"$TERMTOSEARCH"'$/, /^[[:upper:][:punct:]]+$/ {
        # remove unnecessary term pointers
        s/^[[:upper:][:punct:]]+$//
        #remove empty lines
        /^$/ d
        # print the rest from the first section
        p
    }'
# ---------------------------------------------------
# Throw new sed scripts below:
URBANSEDSCRIPT=""
ZIPSEDSCRIPT=""
# ---------------------------------------------------

DICTNAME="$(basename "$DICTFILE")"
DICTNAME="${DICTNAME%%.txt}"

case "$DICTNAME" in
    "webster")
        SEDSCRIPT="$DEFAULT_SEDSCRIPT";;
    # --------------------------------------------
    # Use dedicated sed scripts below:
    "urban")
        SEDSCRIPT="$URBANSEDSCRIPT";;
    "zip")
        SEDSCRIPT="$ZIPSEDSCRIPT";;
    # --------------------------------------------
esac

# ---------------------------------------------------------
DEFINITION=$(fgrep -x -A $MAXCONTEXTLINES "$TERMTOSEARCH" "$DICTFILE")
# using here-strings and moderately advanced sed
RESULT=$(sed -nr "$SEDSCRIPT" <<< "$DEFINITION")
# TODO: make result more intiutive for the user
echo "$RESULT"
# ---------------------------------------------------------

exit $?

# Exercises:
# ---------
# 1)  Modify the script to accept any type of alphabetic input
#   + (uppercase, lowercase, mixed case), and convert it
#   + to an acceptable format for processing.
#
# 2)  Convert the script to a GUI application,
#   + using something like 'gdialog' or 'zenity' . . .
#     The script will then no longer take its argument(s)
#   + from the command-line.
#
# 3)  Modify the script to parse one of the other available
#   + Public Domain Dictionaries, such as the U.S. Census Bureau Gazetteer.

# zenity error widget usage
z_error_term () {
    zenity --error 
}
