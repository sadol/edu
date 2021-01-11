#!/bin/bash
# factr.sh: Factor a number

MIN=2       # Will not work for number smaller than this.
E_NOARGS=85
E_TOOSMALL=86
E_BADARG=87

if [[ -z $1 ]];then
    echo "Usage: $0 number"
    exit $E_NOARGS
fi

if [[ "$1" -lt "$MIN" ]];then
    echo "Number to factor must be $MIN or greater."
    exit $E_TOOSMALL
fi

# remove nondigits
INPUT=$(sed -nr '/[^[:digit:]]//gp' <<< $1)
if [[ ! $INPUT = $1 ]];then
    echo "Argument <$1> must be of integer type."
    exit $E_BADARG
fi

# Exercise: Add type checking (to reject non-integer arg).

echo "Factors of $1:"
# -------------------------------------------------------
dc <<< "$1[p]s2[lip/dli%0=1dvsr]s12sid2%0=13sidvsr[dli%0=1lrli2+dsi!>.]ds.xd1<2"
# -------------------------------------------------------
#  Above code written by Michel Charpentier <charpov@cs.unh.edu>
#  Used in ABS Guide with permission (thanks!).

exit $?

 # $ sh factr.sh 270138
 # 2
 # 3
 # 11
 # 4093
