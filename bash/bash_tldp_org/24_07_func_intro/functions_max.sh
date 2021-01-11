#!/bin/bash
# max.sh: Maximum of two integers.

# return codes for the `max2' function
E_PARAM_ERR=-2                      # If less than 2 params passed to function.
EQUAL=-1                                   # Return value if both params equal.
FIRSTBIGGER=1
SECONDBIGGER=2
NUMPARAM=2
QUIT="q"

# INFO: comparator function for 2 integers
# USAGE : max2 <first> <second>
# RETURNS : -2 → parameter error
#           -1 → `first' == `second'
#            1 → `first' > `second'
#            2 → `first' < `second'
max2() {
    if [[ $# -ne $NUMPARAM ]];then
        echo $E_PARAM_ERR
    else
        if [[ $1 -eq $2 ]];then
            echo $EQUAL
        elif [[ $1 -gt $2 ]];then
            echo $FIRSTBIGGER
        else
            echo $SECONDBIGGER
        fi
    fi
}


echo "----------------------------------"
echo "Welcome to the GREAT COMPARATOR!!!"
echo "----------------------------------"
echo

while :;do
    echo -n "Enter first number to compare (<q> to quit): "
    read first
    [[ "$first" = $QUIT ]] && break
    echo -n "Enter second number to compare (<q> to quit): "
    read second
    [[ "$second" = $QUIT ]] && break

    ret=$(max2 $first $second)
    case $ret in
        $EQUAL) echo "The two numbers are equal.";;
        $FIRSTBIGGER) echo "<$first> is bigger than <$second>.";;
        $SECONDBIGGER) echo "<$first> is smaller than <$second>.";;
        *) echo "You fool, try again.";;
    esac
done

echo "-------------------------------------"
echo "Don't leave me! You will regret this!"
echo "-------------------------------------"
echo

exit $?

#  Exercise (easy):
#  ---------------
#  Convert this to an interactive script,
#+ that is, have the script ask for input (two numbers).
