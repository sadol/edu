#!/bin/bash

# Arabic number to Roman numeral conversion
# It's crude, but it works.

# Usage: roman <number-to-convert>
E_ARG_ERR=65
NUMBER=$1

if [[ $# -ne 1 || -z "$1" ]];then
    echo "Usage: $(basename $0) <positive-int-to-convert>."
    exit $E_ARG_ERR
fi

# maybe long but faster than using `sed' or `awk'
to_roman_long() {
    [[ $NUMBER -lt 0 ]] && NUMBER=$((-NUMBER)) # there is little sense for negative roman numerals
    local numeral=""
    local thousands=$(($1 / 1000))

    if [[ $thousands -ge 1 ]];then
        for (( i=0; i<$thousands; i++ )) {
            numeral="$numeral"M
            ((NUMBER-=1000))
        }
    fi

    # and now `NUMBER' is less than 1000
    local hundrets=$(( NUMBER / 100 ))
    if [[ $hundrets -ge 1 ]];then
        case $hundrets in
            1) numeral="$numeral"C;;
            2) numeral="$numeral"CC;;
            3) numeral="$numeral"CCC;;
            4) numeral="$numeral"CD;;
            5) numeral="$numeral"D;;
            6) numeral="$numeral"DC;;
            7) numeral="$numeral"DCC;;
            8) numeral="$numeral"DCCC;;
            9) numeral="$numeral"DM;;
        esac
        ((NUMBER-=hundrets*100))
    fi

    # and now `NUMBER' is less than 100
    local tens=$(( NUMBER / 10 ))
    if [[ $tens -ge 1 ]];then
        case $tens in
            1) numeral="$numeral"X;;
            2) numeral="$numeral"XX;;
            3) numeral="$numeral"XXX;;
            4) numeral="$numeral"XL;;
            5) numeral="$numeral"L;;
            6) numeral="$numeral"LX;;
            7) numeral="$numeral"LXX;;
            8) numeral="$numeral"LXXX;;
            9) numeral="$numeral"XC;;
        esac
        ((NUMBER-=tens*10))
    fi

    # and now `NUMBER' is less than 10
    if [[ $NUMBER -ge 1 ]];then
        case $NUMBER in
            1) numeral="$numeral"I;;
            2) numeral="$numeral"II;;
            3) numeral="$numeral"III;;
            4) numeral="$numeral"IV;;
            5) numeral="$numeral"V;;
            6) numeral="$numeral"VI;;
            7) numeral="$numeral"VII;;
            8) numeral="$numeral"VIII;;
            9) numeral="$numeral"IX;;
        esac
    fi

    echo "$numeral"
}

echo "$(to_roman_long $1)"

exit $?
