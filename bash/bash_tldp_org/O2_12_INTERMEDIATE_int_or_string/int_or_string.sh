#!/usr/bin/bash
TRUE=0
FALSE=1

is_int() {
    expr "$1" + 0 > /dev/null 2>&1
    [[ $? -eq 0 ]] && return $TRUE
    return $FALSE
}

#---------some testing
is_int "a"
echo
[[ $? -eq $TRUE ]] && echo "\`a' is an int." || echo "\`a' is a string"
