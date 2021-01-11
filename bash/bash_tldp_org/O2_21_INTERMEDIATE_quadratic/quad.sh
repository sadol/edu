#!/usr/bin/bash

E_ARG=56
USAGE="Usage: $(basename $0) <A-coef> <B-coef> <C-coef> .
All coefficients in the form of floats."

if [[ $# -ne 3 ]];then
    echo "$USAGE" >&2
    exit $E_ARG
fi

SUCCESS=0
FAILURE=1

# checks if string can be safely treated as a number
isfloat() {
    local input="${1/./}"
    expr "$input" + 0 > /dev/null 2>&1
    [[ $? -eq 0 ]] && return $SUCCESS
    return $FAILURE
}

E_NOTNUMBER=57
for arg in $@;do
    isfloat "$arg"
    if [[ $? -ne $SUCCESS ]];then
        echo "$USAGE" >&2
        exit $E_NOTNUMBER
    fi
done

A=$1
B=$2
C=$3
equation="${A}x^2 + ${B}x + $C = 0"

if [[ "$(bc <<< "$A == 0")" -eq 1 ]];then
    echo "{A} coefficent is equal to 0. Equation {$equation} does not have solutions."
    exit $SUCCESS
fi

delta="$(bc <<< "scale=6; ($B^2) - (4 * $A * $C)")"
if [[ $(bc <<< "$delta < 0") -eq 1 ]];then
    echo "Equation: {$equation} does not have solution int the realm of REAL numbers."
    exit $SUCCESS
fi

if [[ "$(bc <<< "$delta == 0")" -eq 1 ]];then
    result="$(bc <<< "scale=6; (-${B}) / (2 * $A)")"
    echo "Equation: {$equation} has one solution: {x = ${result}}."
else
    result1="$(bc <<< "scale=6; (-${B} - sqrt(${delta})) / (2 * $A)")"
    result2="$(bc <<< "scale=6; (-${B} + sqrt(${delta})) / (2 * $A)")"
    echo "Equation: {$equation} has two solutions: {x1 = ${result1} , x2 = ${result2}}."
fi

exit $?
