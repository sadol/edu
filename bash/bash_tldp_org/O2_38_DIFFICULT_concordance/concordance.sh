#!/usr/bin/bash

# lists all word occurrences in file along with line numbers in which each word
# occurs.
E_NO_FILE=45
USAGE="Usage: $(basename $0) <file-name> ."
FILE="$1"
if [[ ! -f "$FILE" ]];then
    echo "$USAGE"  >&2
    exit $E_NO_FILE
fi

# indexes are particular words and values are strings of consecutive line numbers
declare -A a_words

# fills associative array of `a_words' with proper keys & values
_fill_array() {
    local word
    local line
    local grep_ret
    local grep_line
    local p_word                           # potential word to add to the array
    local p_val                           # potential value of the certain word

    while read line; do
        for word in $line; do
            p_word="${word//[[:blank:][:punct:]]/}"                  # clearing
            p_val="${a_words[$p_word]}"                              # checking
            if [[ -z $p_val ]]; then
                a_words[$p_word]=""
                grep_ret="$(grep -n -w $word "$FILE")" # as always in BASH, not very efficient
                while read grep_line;do
                    a_words[$p_word]="${a_words[$p_word]} ${grep_line%%:*}"
                done < <(echo "$grep_ret")                 # " " to preserve \n
            fi
        done
    done < "$FILE"
}

# final results presentation interface
_print_results() {
    local word
    local ids
    local sorted_ids="$(sort <(for ids in ${!a_words[@]}; do echo $ids; done))"

    for word in $sorted_ids ; do
        echo "$word: ${a_words[$word]}"
    done
}

main() {
    _fill_array
    _print_results
}

#---------------------------------MAIN-----------------------------------------
main
