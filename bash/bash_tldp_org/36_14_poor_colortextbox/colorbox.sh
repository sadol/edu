#!/bin/bash
# Draw-box.sh: Drawing a box using ASCII characters.

# Script by Stefano Palmeri, with minor editing by document author.
# Minor edits suggested by Jim Angstadt.
# Used in the ABS Guide with permission.


######################################################################
###  draw_box function doc  ###

#  The "draw_box" function lets the user
#+ draw a box in a terminal.
#
#  Usage: draw_box ROW COLUMN HEIGHT WIDTH COLOR "TEXT"
#  ROW and COLUMN represent the position
#+ of the upper left angle of the box you're going to draw.
#  ROW and COLUMN must be greater than 0
#+ and less than current terminal dimension.
#  HEIGHT is the number of rows of the box, and must be > 0.
#  HEIGHT + ROW must be <= than current terminal height.
#  WIDTH is the number of columns of the box and must be > 0.
#  WIDTH + COLUMN must be <= than current terminal width.
#  ${#TEXT} must be shorter than (WIDTH - 2) * (HEIGHT - 2)
#
# E.g.: If your terminal dimension is 20x80,
#  draw_box 2 3 10 45 is good
#  draw_box 2 3 19 45 has bad HEIGHT value (19+2 > 20)
#  draw_box 2 3 18 78 has bad WIDTH value (78+3 > 80)
#
#  COLOR is the color of the box frame.
#  This is the 5th argument and is optional.
#  0=black 1=red 2=green 3=tan 4=blue 5=purple 6=cyan 7=white.
#  If you pass the function bad arguments,
#+ it will just exit with code 65,
#+ and no messages will be printed on stderr.
#
#  Clear the terminal before you start to draw a box.
#  The clear command is not contained within the function.
#  This allows the user to draw multiple boxes, even overlapping ones.

###  end of draw_box function doc  ###
######################################################################

draw_box () {

    HORZ="-"
    VERT="|"
    CORNER_CHAR="+"
    MINARGS=4
    E_BADARGS=65

    [[ $# -lt "$MINARGS" ]] && exit $E_BADARGS

    # check first 5 args: must be numeric
    for ((i=1; i<6; i++)) {
        [[ -n "$(sed -nr '/[^[:digit:]]/ p' <<< ${i})" ]] && exit $E_BADARGS
    }

    BOX_HEIGHT=$(($3 - 1))       #  -1 correction needed because angle char "+"
    BOX_WIDTH=$(($4 - 1))            #+ is a part of both box height and width.
    T_ROWS=$(tput lines)                   #  Define current terminal dimension
    T_COLS=$(tput cols)                                 #+ in rows and columns.
    MAX_LINE=$(( $4 - 2 ))  # max lenght of the one line of text inside the box
    MAX_LEN=$(( ($3 - 2) * MAX_LINE ))          # max length of the text string

    [[ $1 -lt 1  ||  $1 -gt $T_ROWS ]] && exit $E_BADARGS
    [[ $2 -lt 1  ||  $2 -gt $T_COLS ]] && exit $E_BADARGS
    [[ $(($1 + BOX_HEIGHT + 1)) -gt $T_ROWS ]] && exit $E_BADARGS
    [[ $(($2 + BOX_WIDTH + 1)) -gt $T_COLS ]] && exit $E_BADARGS
    [[ $3 -lt 1  ||  $4 -lt 1 ]] && exit $E_BADARGS
    [[ ${#6} -gt $MAX_LEN ]] && exit $E_BADARGS

    plot_char() {                                 # Function within a function.
        echo -e "\E[${1};${2}H"$3  # $3 may be the whole string (not only char)
    }

    echo -ne "\E[3${5}m"                     # Set box frame color, if defined.

    # start drawing the box
    for ((r=$1,count=1,p_text=0;count<=BOX_HEIGHT;r++,count++)) {
        plot_char $r $2 $VERT
        if [[ $r -gt 2 ]];then
            plot_char $r $(($2 + 1)) "${6:$p_text:$MAX_LINE}"  # CRUDE & SIMPLE
            ((p_text+=MAX_LINE))
        fi
    }

    for ((r=$1,count=1,c=$2+BOX_WIDTH;count<=BOX_HEIGHT;r++,count++)) { plot_char $r $c $VERT; }
    for ((c=$2,count=1;count<=BOX_WIDTH;c++,count++)) { plot_char $1 $c $HORZ; }
    for ((c=$2,count=1,r=$1+BOX_HEIGHT;count<=BOX_WIDTH;c++,count++)) { plot_char $r $c $HORZ; }
    plot_char $1 $2 $CORNER_CHAR                             # Draw box angles.
    plot_char $1 $(($2 + BOX_WIDTH)) $CORNER_CHAR
    plot_char $(($1 + BOX_HEIGHT)) $2 $CORNER_CHAR
    plot_char $(($1 + BOX_HEIGHT)) $(($2 + BOX_WIDTH)) $CORNER_CHAR

    echo -ne "\E[0m"                                     #  Restore old colors.
    P_ROWS=$((T_ROWS-10))           #  Put the prompt at bottom of the terminal.
    echo -e "\E[${P_ROWS};1H"
}


# Now, let's try drawing a box.
clear                       # Clear the terminal.
R=2                                                                       # Row
C=3                                                                    # Column
H=10                                                                   # Height
W=45                                                                    # Width
col=1                                                             # Color (red)
song="Putin has a small dick, which Xi wants to lick."                   # text
draw_box $R $C $H $W $col "$song"                               # Draw the box.

exit 0

# Exercise:
# --------
# Add the option of printing text within the drawn box.
