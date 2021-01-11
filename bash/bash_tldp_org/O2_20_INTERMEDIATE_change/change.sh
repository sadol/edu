#!/usr/bin/bash

E_ARG=67
USAGE="Usage: $(basename $0) <float-with-two-precission>.
For example: \`$(basename $0) 12.03'."
CHANGE="$1"
if [[ -z $(sed -nr '/^[[:digit:]]*[,.][[:digit:]]{2}$/p' <<< "$CHANGE") ]];then
    echo "$USAGE" >&2
    exit $E_ARG
fi

# change billon in my country (smallest denominator)
FIFTY=50
TWENTY=20
TEN=10
FIVE=5
ONE=1

# main value times 100 to remove fractions
CHANGE="$(sed -nr 's/[,.]//p' <<< "$CHANGE")"

N50=$(( CHANGE / FIFTY ))
CHANGE=$(( CHANGE % FIFTY ))
N20=$(( CHANGE / TWENTY ))
CHANGE=$(( CHANGE % TWENTY ))
N10=$(( CHANGE / TEN ))
CHANGE=$(( CHANGE % TEN ))
N5=$(( CHANGE / FIVE ))
CHANGE=$(( CHANGE % FIVE ))
N1=$CHANGE

echo "Change of \`$1' is :
    $N50 of 00.${FIFTY},
    $N20 of 00.${TWENTY},
    $N10 of 00.${TEN},
    $N5 of 00.${FIVE} and
    $N1 of 00.0${ONE}.
"

exit $?
