#!/usr/bin/bash

SUCCESS=0
FAILURE=1
# initial search of illegal characters
# INFO: `-' char is placed at the END of the range!!!!!!!!!!!!!!!!!
SEDINIT='/[^[:digit:][:space:],.+-]/ p'
# remove unnecesary characters
SEDREMOVE='s/[[:space:],]+//g; p'
# check the minuses
SEDMINUS='s/^-+//g; p'
# check the pluses
SEDPLUS='s/^\++//g; p'
# check leading zeroes
SEDLEADZERO='s/^0+//; p'
# check if first char is a `.'
SEDLEADDOT='s/^\.+//; p'
# for `0000.00'exception  handling
SEDZEROEXCEPTION='s/0//g; p'
# trailing fraction zeroes in a integer
SEDTRAILZEROES='s/\.0+$//; p'

# INFO: overcomplicated; should be callig some C function instead
# tries to convert string val to integer, echoes back correct int and status
atoi() {
    local integer="$1"
    local minusflag=0                                     # is there any minus?
    local lenpre lenpost

    [[ -n "$(sed -nr "$SEDINIT" <<< "$integer")" ]] && return $FAILURE
    integer="$(sed -nr "$SEDREMOVE" <<< "$integer")"

    if [[ "$integer" =~ ^-.* ]];then
        minusflag=1
        integer=${integer:1}
    fi
    lenpre=${#integer}
    integer="$(sed -nr "$SEDMINUS" <<< "$integer")"
    lenpost=${#integer}
    [[ $(( lenpre - lenpost )) -ne 0 ]] && return $FAILURE

    lenpre=${#integer}
    integer="$(sed -nr "$SEDPLUS" <<< "$integer")"
    lenpost=${#integer}
    [[ $(( lenpre - lenpost )) -gt 1 ]] && return $FAILURE

    integer="$(sed -nr "$SEDLEADZERO" <<< "$integer")"

    # check the rest of the string after removing potential leading zeroes
    if [[ "${integer:0:1}" = "." ]];then
        if [[ "$(sed -nr "$SEDZEROEXCEPTION" <<< "$integer")" != "." ]];then
            return $FAILURE
        else                                      # special case of `000.00000'
            integer="0"; echo $integer
            return $SUCCESS
        fi
    fi

    lenpre=${#integer}
    integer="$(sed -nr "$SEDLEADDOT" <<< "$integer")"
    lenpost=${#integer}
    [[ $(( lenpre - lenpost )) -gt 0 ]] && return $FAILURE # `.98' is illegal then

    integer="$(sed -nr  <<< "$integer")"

    # enough of these pesky checks
    expr $integer + 0 > /dev/null 2>&1
    [[ $? -ne $SUCCESS ]] && return $FAILURE
    [[ $minusflag -eq 1 ]] && integer="-"${integer}
    echo "$integer"
    return $SUCCESS
}

# some tests
first="first "
second=" 2"
third="-3"
fourth="-4 .00"
fifth="005 "
sixth="-06.00"
seventh="- 007.01"
eight="-a008"

test_atoi() {
    atoi "$1"
    [[ $? -eq $SUCCESS ]] && echo "$1 OK" || echo "$1 NOT OK"
}

test_atoi "$first"
test_atoi "$second"
test_atoi "$third"
test_atoi "$fourth"
test_atoi "$fifth"
test_atoi "$sixth"
test_atoi "$seventh"
test_atoi "$eight"
