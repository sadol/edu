#!/bin/bash
# life.sh: "Life in the Slow Lane"
# Author: Mendel Cooper
# License: GPL3

# Version 0.2:   Patched by Daniel Albers
#+               to allow non-square grids as input.
# Version 0.2.1: Added 2-second delay between generations.

# ##################################################################### #
# This is the Bash script version of John Conway's "Game of Life".      #
# "Life" is a simple implementation of cellular automata.               #
# --------------------------------------------------------------------- #
# On a rectangular grid, let each "cell" be either "living" or "dead."  #
# Designate a living cell with a dot, and a dead one with a blank space.#
#      Begin with an arbitrarily drawn dot-and-blank grid,              #
#+     and let this be the starting generation: generation 0.           #
# Determine each successive generation by the following rules:          #
#   1) Each cell has 8 neighbors, the adjoining cells                   #
#+     left, right, top, bottom, and the 4 diagonals.                   #
#                                                                       #
#                       123                                             #
#                       4*5     The * is the cell under consideration.  #
#                       678                                             #
#                                                                       #
# 2) A living cell with either 2 or 3 living neighbors remains alive.   #
SURVIVE=2                                                               #
# 3) A dead cell with 3 living neighbors comes alive, a "birth."        #
BIRTH=3                                                                 #
# 4) All other cases result in a dead cell for the next generation.     #
# ##################################################################### #


# Read the starting generation from the file "gen0" ...
# Default, if no other file specified when invoking script.
startfile=${1:-"gen0"}
#  Abort script if "startfile" not specified or "gen0" not present.
E_NOSTARTFILE=86
if [[ ! -f "$startfile" ]];then
    echo "Startfile \""$startfile"\" missing!"
    exit $E_NOSTARTFILE
fi
############################################


# Represent living and dead cells in the start-up file.
ALIVE1=.
DEAD1=_

#  -----------------------------------------------------#
#  This script uses a 10 x 10 grid (may be increased,
#+ but a large grid will slow down execution).
ROWS=10
COLS=10
#  Change above two variables to match desired grid size.
#  -----------------------------------------------------#

GENERATIONS=10          #  How many generations to cycle through.
                        #  Adjust this upwards
                        #+ if you have time on your hands.

NONE_ALIVE=85           #  Exit status on premature bailout,
                        #+ if no cells left alive.
DELAY=1                 #  Pause between generations.
TRUE=0
FALSE=1
ALIVE=0
DEAD=1

declare -a avar                             # Global; holds current generation.
generation=0                                     # Initialize generation count.
alive=0               # How many cells alive at any given time. Initially zero.

# =================================================================

cells=$(( ROWS * COLS ))                                      # How many cells.

# Arrays containing "cells."
declare -a initial
declare -a current

display () {
    alive=0

    local -a arr=( $(echo "$1") )
    local count_max=${#arr[*]}
    local i rowcheck cell

    for ((i=0; i<count_max; i++)) {
        # Insert newline at end of each row.
        rowcheck=$(( i % COLS ))
        if [[ "$rowcheck" -eq 0 ]];then
            echo                # Newline.
            echo -n "      "    # Indent.
        fi

        cell=${arr[i]}
        [[ "$cell" = "." ]] && (( alive++ ))

        # Print out array, changing underscores to spaces.
        echo -n "${cell//_/ }"
    }
}

IsValid () {                                   # Test if cell coordinate valid.
    [[ -z "$1" || -z "$2" ]] && return $FALSE
    [[ "$1" -lt 0 || "$1" -gt $(( ROWS * COLS - 1 )) ]] && return $FALSE

    local row=$2
    local left=$(( row * COLS ))                                  # Left limit.
    local right=$(( left + COLS - 1 ))                           # Right limit.

    [[ "$1" -lt "$left" || "$1" -gt "$right" ]] && return $FALSE

    return $TRUE                                            # Valid coordinate.
}

#  Test whether cell is alive.
#  Takes array, cell number, and state of cell as arguments.
#  Get alive cell count in neighborhood.
IsAlive () {
    GetCount "$1" $2
    local nhbd=$?

    # Alive in any case.
    [[ "$nhbd" -eq "$BIRTH" ]] && return $ALIVE

    # Alive only if previously alive.
    [[ "$3" = "." && "$nhbd" -eq "$SURVIVE" ]] && return $ALIVE

    # Defaults to dead.
    return $DEAD
}

GetCount () {            # Count live cells in passed cell's neighborhood.
                         # Two arguments needed:
	                		# $1) variable holding array
		                  	# $2) cell number
    local cell_number=$2
    local array=( $(echo "$1") )
    local top=$(( cell_number - COLS - 1 ))         # Set up cell neighborhood.
    local center=$(( cell_number - 1 ))
    local bottom=$(( cell_number + COLS - 1 ))
    local r=$(( cell_number / COLS ))
    local row i t_top t_cen t_bot
    local count=0
    local ROW_NHBD=3

    for ((i=0; i<ROW_NHBD; i++)) {               # Traverse from left to right.
        t_top=$(( top + i ))
        t_cen=$(( center + i ))
        t_bot=$(( bottom + i ))

        row=$r                                              # Count center row.
        IsValid $t_cen $row
        [[ $? -eq $TRUE && ${array[$t_cen]} = $ALIVE1 ]] && (( count++ ))

        row=$(( r - 1 ))                                       # Count top row.
        IsValid $t_top $row
        [[ $? -eq $TRUE && ${array[$t_top]} = $ALIVE1 ]] && (( count++ ))

        row=$(( r + 1 ))                                    # Count bottom row.
        IsValid $t_bot $row
        [[ $? -eq $TRUE && ${array[$t_bot]} = $ALIVE1 ]] && (( count++ ))
    }

    [[ ${array[$cell_number]} = $ALIVE1 ]] && (( count-- ))
    return $count
}


next_gen () {                                        # Update generation array.
    local -a array=( $(echo "$1") )
    local i=0

    for ((i=0; i<cells; i++)) {
        IsAlive "$1" $i ${array[$i]}                       # Is the cell alive?
        [[ $? -eq $ALIVE ]] && array[$i]="." || array[$i]="_"
    }

    #    let "generation += 1"       # Increment generation count.
    ###  Why was the above line commented out?

    # Set variable to pass as parameter to "display" function.
    avar=$(echo ${array[@]})           # Convert array back to string variable.
    display "$avar"                                               # Display it.
    echo; echo
    echo "Generation $generation  -  $alive alive"

    if [[ "$alive" -eq 0 ]];then
        echo
        echo "Premature exit: no more cells alive!"
        exit $NONE_ALIVE
    fi
}

# =========================================================

# main ()
# {

# Load initial array with contents of startup file.
initial=( $(sed -nr '/#/d; s/^(.+)$/& /; s/\./\. /g; s/_/_ /g; p' "$startfile") )

clear

echo #         Title
setterm -reverse on
echo "======================="
setterm -reverse off
echo "    $GENERATIONS generations"
echo "           of"
echo "\"Life in the Slow Lane\""
setterm -reverse on
echo "======================="
setterm -reverse off

sleep $DELAY   # Display "splash screen" for 2 seconds.


# -------- Display first generation. --------
Gen0=$(echo ${initial[@]})
display "$Gen0"           # Display only.
echo; echo
echo "Generation $generation  -  $alive alive"
sleep $DELAY
# -------------------------------------------


(( generation++ ))     # Bump generation count.
echo

# ------- Display second generation. -------
Cur=$(echo ${initial[@]})
next_gen "$Cur"          # Update & display.
sleep $DELAY
# ------------------------------------------

(( generation++ ))     # Bump generation count.

# ------ Main loop for displaying subsequent generations ------
for (( ;generation<GENERATIONS ;generation++ )) {
    Cur="$avar"
    next_gen "$Cur"
    sleep $DELAY
}
# ==============================================================

echo
# }

exit 0   # CEOF:EOF

# The grid in this script has a "boundary problem."
# The the top, bottom, and sides border on a void of dead cells.
# Exercise: Change the script to have the grid wrap around,
# +         so that the left and right sides will "touch,"
# +         as will the top and bottom.
#
# Exercise: Create a new "gen0" file to seed this script.
#           Use a 12 x 16 grid, instead of the original 10 x 10 one.
#           Make the necessary changes to the script,
#+          so it will run with the altered file.
#
# Exercise: Modify this script so that it can determine the grid size
#+          from the "gen0" file, and set any variables necessary
#+          for the script to run.
#           This would make unnecessary any changes to variables
#+          in the script for an altered grid size.
#
# Exercise: Optimize this script.
#           It has redundant code.
