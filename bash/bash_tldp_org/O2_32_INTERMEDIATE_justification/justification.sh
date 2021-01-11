#!/usr/bin/bash

E_FILE=45
E_WIDTH=46
USAGEMSG="Usage: $(basename $0) [-f <file_name>] [-w <width>]"
FILE=""
WIDTH=1                                                           # default w=1

while getopts ":f:w:" OPT;do
    case $OPT in
        f) FILE="$OPTARG" ;;
        w) WIDTH="$OPTARG" ;;
        :)
            echo "Missing argument for option <$OPTARG>." >&2
            echo "$USAGEMSG" >&2
            exit $E_WRONGARG
            ;;
        \?)
            echo "$USAGEMSG" >&2
            exit $E_WRONGARG
            ;;
    esac
done

shift $(( OPTIND - 1 ))

if [[ -n $FILE ]];then
    if [[ ! -f "$FILE" ]];then
        echo "No such file : <"$FILE"> in the filesystem." >&2
        echo "$USAGEMSG" >&2
        exit $E_FILE
    fi
else
    FILE="/dev/stdin"
fi

(( WIDTH + 0 )) 2>/dev/null >&2
integer=$?
minus="$(( WIDTH <= 0 ))"
if [[ $integer -eq 1 || $minus -eq 1 ]];then
    echo "Value of width : <$WIDTH> must be positive integer." >&2
    echo "$USAGEMSG" >&2
    exit $E_WIDTH
fi

# Calculates amount of spaces between respective words in a line.
# INPUT: line of text to be spaced accordingly.
# OUTPUT: properly spaced line of text.
justify() {
    local line="$1"
    local no_chars="${line//[[:space:]]/}"
    no_chars=${#no_chars}                               # no of non-white chars
    local -a words=( $1 )                    # INFO: lack of `"' removes spaces
    local no_words=${#words[@]}
    local -a out_words=( $1 )          # words with appriopriate no of spaces prefixed

    local min_width=0
    local i
    for ((i=0; i<no_words; i++)) {       # add up lenghts of the separate words
        ((min_width += ${#words[i]} ))
    }
    (( min_width += no_words - 1 ))                        # and ad some spaces
    if (( min_width >= WIDTH )); then                # wrong WIDTH has been set
        echo "${out_words[@]}"
        return
    fi

    for ((i=0; i<no_words; i++)) { # now explicitly add new spaces between words
        ((min_width += " ${#words[i]}" ))
    }

    # add some NEW spaces accordingly;
    # but first calculate how much spaces to add to the output lines as a whole
    local spaces_to_add=0
    spaces_to_add=$(( WIDTH - no_words - 1 - no_chars ))
    #                         ^^^^^^^^^^^^   ^^^^^^^^
    #                         spaces         non-spaces
    # distribute said spaces between words
    local no_first_add=$(( spaces_to_add / no_words ))       # how namy spaces to add in the first batch
    local no_rest_add=$(( spaces_to_add % no_words ))        # how many spaces left
    local first_add=""                                       # spaces substring to add to each word as a prefix

    if (( no_first_add > 0 )); then
        for (( i=0; i<no_first_add; i++ )) {
            first_add+=" "
        }
        for (( i=0; i<no_words; i++ )) {
            out_words[i]="${first_add}${out_words[i]}"
        }
    fi

    if (( no_rest_add > 0 )); then
        for (( i=0; i<no_rest_add; i++ )) {     # a little bit of asymetry here
            out_words[i]=" ${out_words[i]}"
        }
    fi

    echo "${out_words[@]}"              # INFO: without `"' output is incorrect
}

#-------------------------------main()-----------------------------------------
while read line; do
    justify "$line"
done < "$FILE"
exit $?
