#!/usr/bin/bash

# primes time!!!
PRIMES=( 2 )                                              # must be initialized
USAGE="Usage: $(basename $0) <star-value> <stop-value>."

E_ARG=89
if [[ $1 -le 0 || $2 -le 0 || $1 -ge $2 ]];then        # very rudimentary check
    echo "$USAGE" >&2
    exit $E_ARG
fi

STARTVAL=$1
STOPVAL=$2

populate_primes() {
    local prime check could_be_prime

    for (( check=${PRIMES[0]} + 1; check<$STOPVAL; check++ )) {  # check all numbers in the given range
        could_be_prime=0
        for prime in ${PRIMES[@]};do                  # against all known primes
            if [[ $(( check / 2 )) -gt $prime ]];then # but not greater than 1/2 of the original value
                if [[ $(( check % prime )) -ne 0 ]];then
                    could_be_prime=1
                    continue
                else
                    could_be_prime=0
                    break
                fi
            else
                break                                  # no need to check bigger values
            fi
        done

        [[ $could_be_prime -eq 1 ]] && PRIMES[ ${#PRIMES[@]} ]=$check
    }
}

print_primes() {
    local i j=0
    echo
    echo "---------primes form \`$STARTVAL' to \`$STOPVAL'--------------"
    echo
    for (( i=0; i<${#PRIMES[@]}; i++ )) {
        if [[ ${PRIMES[i]} -ge $STARTVAL ]];then
            if [[ ${PRIMES[i]} -ge $2 ]];then
                if (( (j + 1) % 5 ));then
                    printf '%s\t' "${PRIMES[i]}"
                else
                    printf '%s\n' "${PRIMES[i]}"
                fi
            fi
            ((j++))
        fi                                                        # fi fi ri fi
    }
    echo
    echo "---------"
}

populate_primes $STARTVAL $STOPVAL
print_primes
