#!/usr/bin/bash

E_MOVE=56                                                          # wrong move
E_USER=57
declare -a table                                                   # game array
TABLEN=9
declare -a next_moves # ids of possible moves for server player(recomendations)
TABSIZE="$(bc <<< "sqrt($TABLEN)")"                        # for bigger tables?
SERVERCHAR="X"
HUMANCHAR="O"                                                         # or 25ef
EMPTYCHAR=" "
BORDER_LINE="--+-+--\n"
BORDER_MID="  | | \n"
BORDERS="${BORDER_MID}${BORDER_LINE}${BORDER_MID}${BORDER_LINE}${BORDER_MID}"
 # translation table to cast `table' position into `BORDERS' position for print
declare -a coords=( 2 4 6 19 21 23 36 38 40 )
x=
y=
HUMANLOSTMSG="You lost, I won."
SERVERLOSTMSG="You won this time."
INITMSG="Do you want to play a game? (y|n): "
AGAINMSG="Do you want to play again? (y|n): "
TURNMSG="Your turn. Enter coordinates (x y): "
DRAWMSG="Abort , abort, abort."
YES="y"
NO="n"
answer=""
TURN=0
WON=0
LOST=0
SERVER=0
HUMAN=1
SUCCESS=0
FAILURE=1
START=1

# initialization of the `table' table table table ...
_init_table() {
    local i
    for ((i=0; i<TABLEN; i++)) {
        table[i]=$EMPTYCHAR
    }
}

# initialization of the `next_moves' table
_init_next_moves() {
    local i
    for ((i=0; i<TABLEN; i++)) {
        next_moves[i]=$EMPTYCHAR
    }
}

E_UPDATE_TABLE=58
# casts users' input coordinates of `x' & `y' into certain place in the `table'
_update_table() {
    local table_position=$(( TABSIZE*x + y  ))

    if [[ $x -gt $TABSIZE || $y -gt $TABSIZE || $x -lt 0 || $y -lt 0 ]];then
        exit $E_UPDATE_TABLE
    else
        if [[ ${table[table_position]} != "$EMPTYCHAR" ]];then
            exit $E_UPDATE_TABLE
        else
            table[table_position]=$HUMANCHAR
        fi
    fi
}

# updates contents of the printed table accordingly BUT does NOT change initial
# `BORDERS' string, generates new string instead.
_update_borders() {
    local new_borders=""                                        # return string
    local i                                       # `for' loop sentinel/counter
    local position                                # position of the next insert

    for ((i=0; i<TABLEN; i++)) {
        case $i in
            0) new_borders="${BORDERS:0:1}${table[i]}${BORDERS:${coords[i]}}";;
            *)
                position=$(( ${coords[i]} - 1 ))
                new_borders="${new_borders:0:$position}${table[i]}${BORDERS:${coords[i]}}"
                ;;
        esac
    }

    echo "$new_borders"
}

# checks if winning  conditions as fullfiled; very naive and not scalable
# version (but simple);
# INPUT: `player_id'
# OUTPUT: `y' | `n'
_has_won() {
    local player="$1"
    local char=""
    local i
    local result=$NO

    [[ $player -eq $SERVER ]] && char="$SERVERCHAR" || char="$HUMANCHAR"
    # horizontal check & vertical check
    for (( i=0; i<TABLEN; i+=TABSIZE )) {
        [[ ${table[i]} != $char ]] && continue
        [[ ${table[i]} == ${table[i+1]} && ${table[i+2]} == ${table[i+1]} ]] && result=$YES
        [[ ${table[i]} == ${table[i+TABSIZE]} &&  ${table[i+(TABSIZE*2)]} == ${table[i+TABSIZE]} ]] && result=$YES
    }

    # diagonal check
    [[ ${table[0]} == $char && ${table[0]} == ${table[4]} && ${table[4]} == ${table[8]} ]] && result=$YES
    [[ ${table[2]} == $char && ${table[2]} == ${table[4]} && ${table[4]} == ${table[6]} ]] && result=$YES

    echo $result
}

# checks if there is any chance to make winnig move
# OUTPUT: recomendation (linear) coordinates to counter the move
_recomend_move() {
    local i
    _init_next_moves

    if [[ $TURN -eq 1 ]];then
        if [[ ${table[4]} == $EMPTY ]];then # best second move (first is always human player)
            next_moves[4]=$SERVERCHAR                           # 'golden' move
        else                                                  # choose randomly
            for ((i=0; i<TABLEN; i++)) {
                [[ ${table[i]} == $EMPTYCHAR ]] && next_moves[i]=$SERVERCHAR
            }
        fi
    elif [[ $TURN -ge 5 ]];then
        # try to use own moves to win (at least sixth move when human player is dumb)
        for ((i=0; i<TABLEN; i+=TABSIZE )) {                           # horizontally
            [[ ${table[i]} == $SERVERCHAR && ${table[i+1]} == $SERVERCHAR && ${table[i+2]} == $EMPTYCHAR ]] && next_moves[i+2]=$SERVERCHAR
            [[ ${table[i+1]} == $SERVERCHAR && ${table[i+2]} == $SERVERCHAR && ${table[i]} == $EMPTYCHAR ]] && next_moves[i]=$SERVERCHAR
            [[ ${table[i]} == $SERVERCHAR && ${table[i+2]} == $SERVERCHAR && ${table[i+1]} == $EMPTYCHAR ]] && next_moves[i+1]=$SERVERCHAR
        }

        for ((i=0; i<TABSIZE; i++ )) {                             # vertically
            [[ ${table[i]} == $SERVERCHAR && ${table[i+3]} == $SERVERCHAR && ${table[i+6]} == $EMPTYCHAR ]] && next_moves[i+6]=$SERVERCHAR
            [[ ${table[i+3]} == $SERVERCHAR && ${table[i+6]} == $SERVERCHAR && ${table[i]} == $EMPTYCHAR ]] && next_moves[i]=$SERVERCHAR
            [[ ${table[i]} == $SERVERCHAR && ${table[i+6]} == $SERVERCHAR && ${table[i+3]} == $EMPTYCHAR ]] && next_moves[i+3]=$SERVERCHAR
        }

                                                                # diagonally
        [[ ${table[0]} == $SERVERCHAR && ${table[4]} == $SERVERCHAR && ${table[8]} == $EMPTYCHAR ]] && next_moves[8]=$SERVERCHAR
        [[ ${table[8]} == $SERVERCHAR && ${table[4]} == $SERVERCHAR && ${table[0]} == $EMPTYCHAR ]] && next_moves[0]=$SERVERCHAR
        [[ ${table[2]} == $SERVERCHAR && ${table[4]} == $SERVERCHAR && ${table[6]} == $EMPTYCHAR ]] && next_moves[6]=$SERVERCHAR
        [[ ${table[6]} == $SERVERCHAR && ${table[4]} == $SERVERCHAR && ${table[2]} == $EMPTYCHAR ]] && next_moves[2]=$SERVERCHAR
    else
        # check rival's moves and act accordingly (at least fourth move)
        for ((i=0; i<TABLEN; i+=TABSIZE )) {                           # horizontally
            [[ ${table[i]} == $HUMANCHAR && ${table[i+1]} == $HUMANCHAR && ${table[i+2]} == $EMPTYCHAR ]] && next_moves[i+2]=$SERVERCHAR
            [[ ${table[i+1]} == $HUMANCHAR && ${table[i+2]} == $HUMANCHAR && ${table[i]} == $EMPTYCHAR ]] && next_moves[i]=$SERVERCHAR
            [[ ${table[i]} == $HUMANCHAR && ${table[i+2]} == $HUMANCHAR && ${table[i+1]} == $EMPTYCHAR ]] && next_moves[i+1]=$SERVERCHAR
        }

        for ((i=0; i<TABSIZE; i++ )) {                             # vertically
            [[ ${table[i]} == $HUMANCHAR && ${table[i+3]} == $HUMANCHAR && ${table[i+6]} == $EMPTYCHAR ]] && next_moves[i+6]=$SERVERCHAR
            [[ ${table[i+3]} == $HUMANCHAR && ${table[i+6]} == $HUMANCHAR && ${table[i]} == $EMPTYCHAR ]] && next_moves[i]=$SERVERCHAR
            [[ ${table[i]} == $HUMANCHAR && ${table[i+6]} == $HUMANCHAR && ${table[i+3]} == $EMPTYCHAR ]] && next_moves[i+3]=$SERVERCHAR
        }

                                                                # diagonally
        [[ ${table[0]} == $HUMANCHAR && ${table[4]} == $HUMANCHAR && ${table[8]} == $EMPTYCHAR ]] && next_moves[8]=$SERVERCHAR
        [[ ${table[8]} == $HUMANCHAR && ${table[4]} == $HUMANCHAR && ${table[0]} == $EMPTYCHAR ]] && next_moves[0]=$SERVERCHAR
        [[ ${table[2]} == $HUMANCHAR && ${table[4]} == $HUMANCHAR && ${table[6]} == $EMPTYCHAR ]] && next_moves[6]=$SERVERCHAR
        [[ ${table[6]} == $HUMANCHAR && ${table[4]} == $HUMANCHAR && ${table[2]} == $EMPTYCHAR ]] && next_moves[2]=$SERVERCHAR
    fi
    # DEBUG
    #echo "next_move: ${next_moves[@]}"
}

print_borders() {
    echo
    printf "$(_update_borders)"
    echo
}

E_NO_CONTI=61
# INFO: function uses globals!
print_message() {
    local msg=""
    local conti

    if [[ $START -eq 1 ]];then
        msg="${msg}${INITMSG}"
        printf "$msg"
        START=0
        read conti
        [[ $conti == $NO ]] && exit $E_NO_CONTI || _init_table
    elif [[ $WON -eq 1 || $LOST -eq 1 ]];then
        if [[ $WON -eq 1 ]];then
            msg="${msg}${HUMANLOSTMSG}\n"
            WON=0
        fi
        if [[ $LOST -eq 1 ]];then
            msg="${msg}${SERVERLOSTMSG}\n"
            LOST=0
        fi
        msg="${msg}${AGAINMSG}"
        printf "$msg"
        read conti
        [[ $conti == $NO ]] && exit $E_NO_CONTI || _init_table
    else
        msg="${msg}${TURNMSG}"
        printf "$msg"
        read x y                                                          # globals
        _update_table $HUMANCHAR
    fi
}

# checks for draw
_check_draw() {
    local answer=$NO
    local server_wins=$NO
    local human_wins=$NO

    # i don't even bother if it's even  neccesary
    server_wins=$(_has_won $SERVER)
    human_wins=$(_has_won $HUMAN)

    [[ $server_wins == $NO && $human_wins == $NO && $TURN -eq 6 ]] && answer=$YES

    echo $answer
}

# makes move using recomendation table `next_move' & changes globals
# accordingly
make_move() {
    local draw=$(_check_draw)
    if [[ $draw == $YES ]];then                                            # draw
        echo $DRAWMSG
        exit 13
    fi
    print_message
    if [[ "$(_has_won $HUMAN)" == $YES ]];then      # (re)set globals here
        WON=0
        LOST=1
        TURN=0
    fi
    _recomend_move
    # use recomendataions somehow
    local i
    for ((i=0; i<TABLEN; i++)) {
        if [[ ${next_moves[i]} == $SERVERCHAR ]];then
            table[i]=${next_moves[i]}
            break  # take a first recomendation
        fi
    }
    if [[ "$(_has_won $SERVER)" == $YES ]];then
        WON=1
        LOST=0
        TURN=0
    fi
    print_borders                                             # updated borders
    ((TURN++))
}

# main()-----------------------------------------------------------------------
clear
while : ;do                                           # event loop of some kind
    make_move
done
#------------------------------------------------------------------------------

exit $?
