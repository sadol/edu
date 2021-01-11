#!/bin/bash
# poem.sh: Pretty-prints one of the ABS Guide author's favorite poems.
TEXT_SPACING="     "
AUTHOR_SPACING=$TEXT_SPACING$TEXT_SPACING
E_FILE=89

[[ ! -f "$1" ]] && echo "No such poem: <$1>." && exit $E_FILE

echo

tput bold   # Bold print.
fmt=

# whole loop below could be easly replaced by `sed's triliner
while IFS= read -r line; do  # `IFS' and `-r' to force `read' to not cut leading spaces
    if [[ ${#line} && "$(sed -nr '/^[[:space:]]+.+$/ p' <<< "$line")"  ]];then
        fmt="$AUTHOR_SPACING"'%s\n'
    elif [[ ${#line} && "$(sed -nr '/^[^[:space:]]+.+$/ p' <<< "$line")"  ]];then
        fmt="$TEXT_SPACING"'%s\n'
    else                                                          # empty lines
        fmt='%s\n'
    fi
    printf "$fmt" "$line"
done < "$1"

tput sgr0   # Reset terminal.
            # See 'tput' docs.
echo

exit $?

# Exercise:
# --------
# Modify this script to pretty-print a poem from a text data file.
