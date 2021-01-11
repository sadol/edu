#!/usr/bin/bash

# callback function for `mapfile' builtin
# INPUT: line of chars
# OUTPUT: reversed line of chars
reverse () {
    local output=""
    local input="$1"
    local i

    for (( i=0; i<${#input}; i++ )) {
        output="${output}${input:${#input}-$i-1:1}"
    }

    echo "$output"
}

mapfile -t LINES < $0

for ((i=${#LINES[@]}-1; i>=0; i--)) {
    reverse "${LINES[$i]}"
}
