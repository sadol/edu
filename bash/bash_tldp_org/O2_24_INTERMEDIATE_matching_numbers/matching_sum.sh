#!/usr/bin/bash

# a little fun with combinatorics (OR NOT, this is only simulation, not
# algebra)

SET="456"
MIN=10000
MAX=99999
result=0

for ((i=MIN; i<=MAX; i++)) {
    temp="${i//[$SET]/}"
    [[ ${#temp} -eq 3 ]] && (( result += i ))
}

echo "TOTAL SUM: $result"
