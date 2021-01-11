#!/bin/bash
# twodim.sh: Simulating a two-dimensional array.

# A one-dimensional array consists of a single row.
# A two-dimensional array stores rows sequentially.

# user defined array
Rows=${1:-5}
Columns=${2:-5}
declare -a alpha     # char alpha [Rows] [Columns];
                     # Unnecessary declaration. Why?
# buffer of spaces to use for tilting purposes instead of dedicated function
# (UGLY HACK!!!)
SPACES="                                                                      "
# list of symbols ready to use for array's graphical representation
declare -a symbols=( 'A' 'Ą' 'B' 'C' 'Ć' 'D' 'E' 'Ę' 'F' 'G' 'H' 'I' 'J' 'K'
'L' 'Ł' 'M' 'N' 'Ń' 'O' 'Ó' 'P' 'Q' 'R' 'S' 'Ś' 'T' 'U' 'W' 'X' 'Y' 'Z' 'Ż' 'Ź'
'a' 'ą' 'b' 'c' 'ć' 'd' 'e' 'ę' 'f' 'g' 'h' 'i' 'j' 'k' 'l' 'ł' 'm' 'n' 'o' 'ó'
'p' 'q' 'r' 's' 'ś' 't' 'u' 'w' 'x' 'y' 'z' 'ż' 'ź' 'π' 'Ω' 'œ' 'Œ' '©' '®' 'ß'
'™' '←' '¥' '↓' '→' 'þ' '↔' 'Þ' 'ð' 'Ð' 'æ' 'Æ' 'ŋ' 'Ŋ' '•' 'ə' 'Ə' 'µ' '∞')

E_ARG=66
# check user supplied values of number of rows and cols
if [[ -z "$(sed -nr '/^[[:digit:]]+$/ p' <<< $Rows)" ||
        -z "$(sed -nr '/^[[:digit:]]+$/ p' <<< $Rows)" ]];then
    echo "Script arguments must be numeric (ints > 0)."
    exit $E_ARG
fi

# max size of the user defined table size=no_rows * no_cols
MAX_SIZE=${#symbols[@]}
size=$((Rows * Columns))
E_SIZE=67
if [[ $size > $MAX_SIZE ]];then
    echo "Table is too big. Send smaller values to the script."
    exit $E_SIZE
fi

# newer version fully linearized
load_alpha() {
    for ((i=0; i<size; i++ )) {
        alpha[$i]="${symbols[$i]}"
    }
}


print_alpha () {
    echo
    for ((i=0; i<size; i++)){
        [[ $((i % Columns)) -eq 0 ]] && echo
        echo -n " ${alpha[$i]}"
    }
    echo
    # The simpler equivalent is
    #     echo ${alpha[*]} | xargs -n $Columns
}


filter () {     # Filter out negative array indices.
    echo -n ".|"  # Provides the tilt.
                  # Explain how.
    if [[ "$1" -ge 0 &&  "$1" -lt "$Rows" && "$2" -ge 0 && "$2" -lt "$Columns" ]];then
        echo -n "${alpha[$(($1 * $Rows + $2))]}|"   #      alpha[$row][$column]
    fi
}

rotate () { #  Rotate the array 45 degrees --
            #+ "balance" it on its lower lefthand corner.
    for ((row=Rows + 1; row > -Rows - 1; row-- )) {
        for ((column=0; column < Columns; column++ )) {
            if [[ "$row" -ge 0 ]];then
                t1=$((column - row))
                t2=$column
            else
                t1=$column
                t2=$((column + row))
            fi
            filter $t1 $t2   # Filter out negative array indices.
                             # What happens if you don't do this?
        }
        echo; echo
    }
    #  Array rotation inspired by examples (pp. 143-146) in
    #+ "Advanced C Programming on the IBM PC," by Herbert Mayer
    #+ (see bibliography).
    #  This just goes to show that much of what can be done in C
    #+ can also be done in shell scripting.
}

# INFO: this function tries to tilt any rectangle according to this reciepe:
#       let's assume we have such a rectangle (3 rows by 2 columns):
#
#       ab
#       cd
#       ef
#
#       Every tilted row has maximum `Min(Rows,Columns)' elements and there are
#       `Rows+Columns-1' such tilted rows; starting point of every tilted row
#       is placed in the first row, from right to left:
#
#     4321
#       cd
#       ef
#
#       This particular example has 4 tilted rows as follows:
#       `b' `ad' `cf' `e' printed in this order.
new_rotate () {
    local -i c               # pointer to the starting column (may be negative)
    local -i id                     # position of the char in the `alpha' table
    local -i max_len=$(($Columns>$Rows?$Rows:$Columns)) # max len of the tilted row

    echo
    for ((k = Columns + Rows - 1; k > 0; k--)) { # how many tilted rows to print
        for ((i = 0; i < max_len ;i++)) { # build legal tilted row char by char
            echo -n ' '                                            # lower tilt
            c=$((k - Rows + i))
            id=$((i * Columns + c))
            if [[ $i -ge $Rows || $c -ge $Columns || $i -lt 0 || $c -lt 0 ]];then
                continue
            else
                if [[ $i -eq 0 && $c -ge 1 ]];then # upper tilt (UGLY HACK!!!)
                    echo -n "${SPACES:0:$(( 2*(k - max_len) ))}"
                fi
                echo -n " ${alpha[$id]}"
            fi
        }
        echo                                                   # new tilted row
    }
}

#--------------- Now, let the show begin. ------------#
load_alpha     # Load the array.
print_alpha    # Print it out.
new_rotate         # Rotate it 45 degrees counterclockwise.
#-----------------------------------------------------#

exit 0

# This is a rather contrived, not to mention inelegant simulation.

# Exercises:
# ---------
# 1)  Rewrite the array loading and printing functions
#     in a more intuitive and less kludgy fashion.
#
# 2)  Figure out how the array rotation functions work.
#     Hint: think about the implications of backwards-indexing an array.
#
# 3)  Rewrite this script to handle a non-square array,
#     such as a 6 X 4 one.
#     Try to minimize "distortion" when the array is rotated.
