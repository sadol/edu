#!/usr/bin/bash

# "meaninful" use of CONTINUE N loop control construct
EXIT_WORD="drop_the_rest_of_the_line"

echo
echo "Some preprocessing..."
echo

echo "Some processing:"
echo "----------------"
while read line; do
    for word in $line ;do
        if [[ "$word" = "$EXIT_WORD" ]]; then
            echo
            continue 2                                         # got outer loop
        fi
        echo -n "$word "
    done
    echo
done < $1

echo "----------------"
echo
echo "Some postprocessing..."
echo
exit
