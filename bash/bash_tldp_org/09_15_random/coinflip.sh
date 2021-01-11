#!/bin/bash
# How random is RANDOM?

RANDOM=$$       # Reseed the random number generator using script process ID.

SIDES=2         # A coin has 2 faces.
MAXFLIPS=1000   # Increase this if you have nothing better to do with your time.
flip=0          # Number of times the coin have been tossed .

heads=0          #  Must initialize counts to zero,
tails=0          #+ since an uninitialized variable is null, NOT zero.

print_result ()
{
    echo
    echo "heads =   $heads"
    echo "tails =   $tails"
    echo
}

update_count()
{
    case "$1" in
        0) ((heads++));;
        1) ((tails++));;
    esac
}

echo

for ((i=0; i<MAXFLIPS; i++)){
    coin1=$((RANDOM % SIDES))
    update_count $coin1
    ((flip++))
}

print_result

exit $?

#  The scores should distribute evenly, assuming RANDOM is random.
#  With $MAXTHROWS at 600, all should cluster around 100,
#+ plus-or-minus 20 or so.
#
#  Keep in mind that RANDOM is a ***pseudorandom*** generator,
#+ and not a spectacularly good one at that.

#  Randomness is a deep and complex subject.
#  Sufficiently long "random" sequences may exhibit
#+ chaotic and other "non-random" behavior.

# Exercise (easy):
# ---------------
# Rewrite this script to flip a coin 1000 times.
# Choices are "HEADS" and "TAILS."
