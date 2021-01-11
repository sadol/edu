#!/bin/bash
# hypotenuse.sh: Returns the "hypotenuse" of a right triangle.
#                (square root of sum of squares of the "legs")

ARGS=2                # Script needs sides of triangle passed.
E_BADARGS=85          # Wrong number of arguments.

if [[ $# -ne "$ARGS" ]];then # Test number of arguments to script.
    echo "Usage: $(basename $0) side_1 side_2" 1>&2
    exit $E_BADARGS
fi


AWKSCRIPT=' { printf( "%3.7f\n", sqrt($1*$1 + $2*$2) ) } '
#             command(s) / parameters passed to awk

# Now, pipe the parameters to awk.
echo -n "Hypotenuse of $1 and $2 = "
#awk "$AWKSCRIPT" <<< "$1 $2"
bcresult=
# not nice to be forced to use interactive subcommand with here-document
results () {
    bcresult=$(bc -l << EOF
               scale=9
               sqrt ( $1 * $1 + $2 * $2 )
EOF
              )
}

results $1 $2
echo ${bcresult}
exit

# Exercise: Rewrite this script using 'bc' rather than awk.
#           Which method is more intuitive?
