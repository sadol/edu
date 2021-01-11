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
CONTINUE=""

NODICTDIRMSG="Dictionaries directory <${DICTDIR}> is not present.
          Directory should contain subdirs with text based dictionary and sed
          script to operate on said dictionary."

# check dictdir first
if [[ ! -d ${DICTDIR} ]];then
    zenity --error --text "$NODICTDIRMSG"
    exit $E_NODICTDIR
fi

zenity --info --text "Welcome to the shitty dict scrpit."

for (( ; ; )) {
    DICTNAME=$(zenity --list  --radiolist --column "Selection" \
    --column "Dictionary" TRUE "webster" FALSE "zip" FALSE "urban")

    TERMTOSEARCH="$(zenity --entry --text 'Please enter a term to search for:' )"
    TERMTOSEARCH="$(tr [[:lower:]] [[:upper:]] <<< $TERMTOSEARCH)"
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

    # webster sed script with argument
    case "$DICTNAME" in
        "webster")
            SEDSCRIPT="$DEFAULT_SEDSCRIPT"
            DICTFILE="$DEFAULT_DICTFILE"
            ;;
        # --------------------------------------------
        # Use dedicated sed scripts below:
        "urban")
            SEDSCRIPT="$URBANSEDSCRIPT"
            DICTFILE="stub"
            ;;
        "zip")
            SEDSCRIPT="$ZIPSEDSCRIPT"
            DICTFILE="stub"
            ;;
        # --------------------------------------------
    esac

    if [[ -z $(sed -n '/[[:upper:][:punct:]]/ p' <<< ${TERMTOSEARCH}) ]];then
        zenity --error --text "Term <$TERMTOSEARCH> is invalid.
        Only valid terms contains [:alpha:] || [:punct:] chars only!"
    else
        DEFINITION="$(fgrep -x -A $MAXCONTEXTLINES "$TERMTOSEARCH" "$DICTFILE")"
        if [[ -z "$DEFINITION" ]];then
            NOTFOUNDMSG="There is no term <${TERMTOSEARCH}> in dictionary <${DICTFILE}>."
            zenity --info --text "$NOTFOUNDMSG"
        else
            # using here-strings and moderately advanced sed
            RESULT="$(sed -nr "$SEDSCRIPT" <<< "$DEFINITION")"
            zenity --info --width=80 --text "$RESULT"
        fi
    fi

    CONTINUE=$(zenity --question --text "Do you want to quit?")
    [[ ! $CONTINUE ]] && break
}

zenity --info --text "Bye."

# ---------------------------------------------------
# Throw new sed scripts below:
URBANSEDSCRIPT=""
ZIPSEDSCRIPT=""
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
