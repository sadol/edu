PASSES=10            #  Number of particle interactions / marbles.
CURRENTPASS=0
ROWS=10               #  Number of "collisions" (or horiz. peg rows).
RANGE=3               #  0 - 2 output range from $RANDOM.
POS=0                 #  Left/right position.
RANDOM=$$             #  Seeds the random number generator from PID
                      #+ of script.
URANDOM=0             # stores the last byte from /dev/urandom
SLOTSREADY=0          # flag to check if `Slots' array is ready to use : fucking spagetti code

declare -a Slots      # Array holding cumulative results of passes.
declare -A Histogram  # Array holding print artefacts for histogram presentation
declare -A Animation  # Array holding animation artifacts
NUMSLOTS=21           # Number of slots at bottom of board.
HISTROWS=0            # Number of rows in histogram 2D array
HISTCOLS=$NUMSLOTS    # NUmber of cols in histogram 2D array
MAXVAL=0              # `MAXVAL' is maximum value stored in some slot
CLUSTER=10            # no. of observations to create one cluster for use in histograms
HISTCHAR="I"          # the character of the histogram stack
HISTEMPTY="."
ANIMMARBLE="o"
ECLUSTER=100

# WARNING: it is sloooooooooooooooooooooooooooooooooooow
Read_Urandom () { # stores the last /dev/urandom byte
    URANDOM="$(od -vAn -N1 -tu4 < /dev/urandom)"
}


Initialize_Animation () {
    for ((row=1; row<=ROWS; row++)) {
        for ((col=1; col<=$NUMSLOTS; col++)) {
            Animation[$row,$col]=$HISTEMPTY
        }
    }
}


Show_Animation () {
    echo "   Crappy animation no. $CURRENTPASS:"; echo
    for ((row=1; row<=ROWS; row++)) {
        for ((col=1; col<=NUMSLOTS; col++)) {
           echo -n "${Animation[$row,$col]}"
        }
        echo
    }
}


Initialize_Slots () { # Zero out all elements of the array.
    for ((i=1; i<=NUMSLOTS; i++ )) {
        Slots[$i]=0
    }

    SLOTSREADY=0
    echo                  # Blank line at beginning of run.
}


Initialize_Histogram () { # ` ' out all elements of the array
    # phase 1. determine the size of the Histogram 2D array
    local tempmax=0
    if [[ $SLOTSREADY -eq 1 ]];then  # you may search for max val now
        for ((i=1; i<=NUMSLOTS; i++ )) {
            ((Slots[$i] > tempmax)) && tempmax="${Slots[$i]}"
        }
    fi

    # ERROR
    if [[ tempmax -le CLUSTER ]];then
        echo "tempmax <= CLUSTER"
        exit $ECLUSTER
    fi

    (( HISTROWS = tempmax / CLUSTER ))

    # INFO: Bash arrays are 1 based!!!
    # phase 2. declare 2D histogram array
    for ((row=1; row<=HISTROWS; row++)) {
        for ((col=1; col<=HISTCOLS; col++)) {
            Histogram[$row,$col]=$HISTEMPTY
        }
    }
}


Fill_Histogram () {
    Initialize_Histogram

    for ((col=1; col<=HISTCOLS; col++ )) {
        local tempRow=$HISTROWS
        # every `CLUSTER' hits prints `I`, the rest is cut off
        for (( v=Slots[col]; v>=CLUSTER; v-=CLUSTER )) {
            Histogram[$tempRow,$col]="$HISTCHAR"
            (( tempRow-- ))
        }
    }
}


Show_Histogram () {
    Fill_Histogram
    echo "   Crappy histogram"; echo
    for ((row=1; row<=HISTROWS; row++)) {
        for ((col=1; col<=HISTCOLS; col++)) {
           echo -n "${Histogram[$row,$col]}"
       }
        echo
    }
}


Show_Slots () {
    echo; echo
    echo -n " "
    for ((i=1; i<=NUMSLOTS; i++)) {   # Pretty-print array elements.
        printf "%3d" ${Slots[$i]}   # Allot three spaces per result.
    }

    echo # Row of slots:
    echo " |__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|"
    echo "                                ||"
    echo #  Note that if the count within any particular slot exceeds 99,
        #+ it messes up the display.
        #  Running only(!) 500 passes usually avoids this.
}


Move () {              # Move one unit right / left, or stay put.
    # move=$RANDOM         # How random is $RANDOM? Well, let's see ...
    Read_Urandom
    move=$URANDOM      # it more uniform spread
    ((move %= RANGE))  # Normalize into range of 0 - 2.
    case "$move" in
        0 ) ;;                   # Do nothing, i.e., stay in place.
        1 ) ((POS--));;          # Left.
        2 ) ((POS++));;          # Right.
        * ) echo -n "Error ";;   # Anomaly! (Should never occur.)
    esac
}


Play () {                    # Single pass (inner loop).
    SHIFT=11                     # Why 11, and not 10?
    ((CURRENTPASS++))
    Initialize_Animation
    for ((i=1; i<=ROWS; i++ )){
        tput cup 1 1                  # move cursor to the upper left corner
        Move                              # calculate marble position
        Animation[$i,$((POS + SHIFT))]=$ANIMMARBLE     # update marble position
        Show_Animation                    # show marble position
        sleep 0.05s
    }
    (( POS += $SHIFT ))          # Shift "zero position" to center.
    (( Slots[$POS]++ ))          # DEBUG: echo $POS
    SLOTSREADY=1
    # echo -n "$POS "
}


Run () {                     # Outer loop.
    tput civis               # hide cursor
    for ((p=0; p<PASSES; p++ )) {
        Play
        POS=0                      # Reset to zero. Why?
    }
    tput cnorm               # show cursor
}


# --------------
# main ()
clear
Initialize_Slots
Run
Show_Animation
clear
Show_Slots
Show_Histogram
# --------------

exit $?

#  Exercises:
#  ---------
#  1) Show the results in a vertical bar graph, or as an alternative,
#+    a scattergram.
#  2) Alter the script to use /dev/urandom instead of $RANDOM.
#     Will this make the results more random?
#  3) Provide some sort of "animation" or graphic output
#     for each marble played.
