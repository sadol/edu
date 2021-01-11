#!/usr/bin/bash
# greatest common divisor by Euclides

ARGS=2
E_AMOUNTARGS=85
E_INTARG=86
E_MSG="Usage: $(basename $0) <first_int> <second_int>."

if [[ $# -ne $ARGS ]]; then
    echo $E_MSG
    exit $E_AMOUNTARGS
fi

if [[ ! $1 =~ ^[0-9]+$ || ! $2 =~ ^[0-9]+$ ]] ; then
    echo $E_MSG
    exit $E_INTARG
fi

gcd() {
    local dividend=$1
    local divisor=$2
    local reminder=1  # must be init to work properly

    until [[ $reminder -eq 0 ]];do
        reminder=$(( dividend % divisor ))
        dividend=$divisor
        divisor=$reminder
    done

    echo "GCD of $1 and $2 is $dividend"
}

gcd $1 $2
