#!/usr/bin/bash

E_BC=45
which bc > /dev/null 2>&1
if [[ $? -eq 1 ]];then
    echo "No \`bc' installed on this system."
    exit $E_BC
fi

START=0
STOP=100
STEP="0.01"

bc_program='scale=6;
for (i='"$START"'; i<='"$STOP"'; i+='"$STEP"'){
print "ln(", i, ") = ", l(i), "\n"
}'

bc -l <<< "$bc_program"

exit $?
