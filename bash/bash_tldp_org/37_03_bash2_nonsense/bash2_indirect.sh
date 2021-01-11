#!/bin/bash
# resistor-inventory.sh
# Simple database / table-lookup application.

E_DATABASE=78
if [ -z "$1" ];then
    echo "Usage: `basename $0` <database-file>" 2>&1
    exit $E_DATABASE
fi
# ============================================================== #
# Data

B1723_value=470                                   # Ohms
B1723_powerdissip=.25                             # Watts
B1723_colorcode="yellow-violet-brown"             # Color bands
B1723_loc=173                                     # Where they are
B1723_inventory=78                                # How many

B1724_value=1000
B1724_powerdissip=.25
B1724_colorcode="brown-black-red"
B1724_loc=24N
B1724_inventory=243

B1725_value=10000
B1725_powerdissip=.125
B1725_colorcode="brown-black-orange"
B1725_loc=24N
B1725_inventory=89

# ============================================================== #

Usercat=""
FOUND=0
echo
echo -n 'Enter catalog number: '
read Usercat

# reading from a file into the set of dedicated variables
while IFS="|" read -r Cat Inv Val Pdissip Loc Ccode;do
    [ "$Cat" != "$Usercat" ] && continue
    FOUND=1
    echo
    echo "Catalog number $Cat:"
    # Now, retrieve value, using indirect referencing.
    echo "There are ${Inv} of  [${Val} ohm / ${Pdissip} watt]\
    resistors in stock."  #        ^             ^
    # As of Bash 4.2, you can replace "ohm" with \u2126 (using echo -e).
    echo "These are located in bin # ${Loc}."
    echo "Their color code is \"${Ccode}\"."
    break
done < "$1"

[ $FOUND -eq 0 ] && echo "Catalog number : $Usercat not found!"

echo; echo

FOUND=0
arr=
# reading from a file into the array
while IFS="|" read -r Cat Inv Val Pdissip Loc Ccode;do
    unset arr
    arr[0]="$Cat"
    arr[1]="$Inv"
    arr[2]="$Val"
    arr[3]="$Pdissip"
    arr[4]="$Loc"
    arr[5]="$Ccode"
    [ "${arr[0]}" != "$Usercat" ] && continue
    FOUND=1
    echo
    echo "Catalog number ${arr[0]}:"
    # Now, retrieve value, using indirect referencing.
    echo "There are ${arr[1]} of  [${arr[2]} ohm / ${arr[3]} watt]\
    resistors in stock."  #        ^             ^
    # As of Bash 4.2, you can replace "ohm" with \u2126 (using echo -e).
    echo "These are located in bin # ${arr[4]}."
    echo "Their color code is \"${arr[5]}\"."
    break
done < "$1"

# Exercises:
# ---------
# 1) Rewrite this script to read its data from an external file.
# 2) Rewrite this script to use arrays,
#+   rather than indirect variable referencing.
#    Which method is more straightforward and intuitive?
#    Which method is easier to code?


# Notes:
# -----
#  Shell scripts are inappropriate for anything except the most simple
#+ database applications, and even then it involves workarounds and kludges.
#  Much better is to use a language with native support for data structures,
#+ such as C++ or Java (or even Perl).

exit 0
