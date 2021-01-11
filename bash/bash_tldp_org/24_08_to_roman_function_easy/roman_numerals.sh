#!/bin/bash

# Arabic number to Roman numeral conversion
# Range: 0 - 3000
# It's crude, but it works.

# Extending the range and otherwise improving the script is left as an exercise.

# Usage: roman <number-to-convert>
LIMIT=3400
E_ARG_ERR=65
NUMBER=$1

if [[ $# -ne 1 || -z "$1" || $1 -gt $LIMIT ]];then
    echo "Usage: $(basename $0) <positive-int-to-convert-less-than-3400>."
    exit $E_ARG_ERR
fi

to_roman () {  # Must declare function before first call to it.
    factor=$1
    rchar=$2
    remainder=$((NUMBER- factor))

    while [[ "$remainder" -ge 0 ]];do
        echo -n $rchar
        ((NUMBER -= factor))
        remainder=$((NUMBER - factor))
    done
    # Exercises:
    # ---------
    # 1) Explain how this function works.
    #    Hint: division by successive subtraction.
    # 2) Extend to range of the function.
    #    Hint: use "echo" and command-substitution capture.
}

to_roman 1000 M
to_roman 900 CM
to_roman 500 D
to_roman 400 CD
to_roman 100 C
to_roman 90 LXXXX
to_roman 50 L
to_roman 40 XL
to_roman 10 X
to_roman 9 IX
to_roman 5 V
to_roman 4 IV
to_roman 1 I

# Successive calls to conversion function!
# Is this really necessary??? Can it be simplified?
echo

exit
