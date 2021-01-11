#!/usr/bin/bash

# STRIPPER - Trump's best friend, strips all comments from bash script.

DEFAULT_FILE="input.sh"
USAGE="USAGE: $(basename $0) [file-name]"
E_NO_FILE=45
E_NO_DEFAULT_FILE=46

if [[ ! -f "$DEFAULT_FILE" ]];then
    echo "Default file: <$DEFAULT_FILE> not found. " >&2
    echo "$USAGE" >&2
    exit $E_NO_DEFAULT_FILE
fi

if [[ -n "$1" ]];then
    if [[ ! -f "$1" ]];then
        echo "$USAGE" >&2
        exit $E_NO_FILE
    else
        DEFAULT_FILE="$1"
    fi
fi

sed -nr 's/#[^!].*//; s/^[[:space:]]*#[[:space:]]*$//; p' "$DEFAULT_FILE"

exit $?
