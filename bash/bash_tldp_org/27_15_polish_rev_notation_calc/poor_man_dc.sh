#!/bin/bash
# stack.sh: push-down stack simulation
#  Similar to the CPU stack, a push-down stack stores data items
#+ sequentially, but releases them in reverse order, last-in first-out.
BP=100            #  Base Pointer of stack array.
                  #  Begin at element 100.

SP=$BP            #  Stack Pointer.
                  #  Initialize it to "base" (bottom) of stack.
declare -a stack
declare -a Data   #  Contents of stack location.
                  #  Must use global variable,
                  #+ because of limitation on function return range.


                  # 100     Base pointer       <-- Base Pointer
                  #  99     First data item
                  #  98     Second data item
                  # ...     More data
                  #         Last data item     <-- Stack pointer

E_PUSHZERO=45
E_PUSHNULL=46
E_SPERROR=47

push () {
    if [[ $# -ne 0 ]];then
        for ARG in $@;do
            if [[ -n "$ARG" ]];then
                if [[ $SP -ge 0 && $SP -le $BP ]];then
                    stack[$SP]="$ARG"
                    ((SP -= 1))
                else
                    exit $E_SPERROR
                fi
            else
                exit $E_PUSHNULL
            fi
        done
    else
        exit $E_PUSHZERO
    fi
}

E_POP=48

pop () {
    HOW_MANY=${1:-1}

    if [[ $SP -lt $BP && $HOW_MANY -gt 0 && $HOW_MANY -le $((BP - SP)) ]];then
        Data=( ${stack[@]:$((SP + 1)):$HOW_MANY} )
        ((SP += HOW_MANY))
    else
        exit $E_POP
    fi
}

status_report () {
    echo "-------------------------------------"
    echo "REPORT"
    echo "Stack Pointer = $SP"
    echo "Just popped \""${Data[@]}"\" off the stack."
    echo "-------------------------------------"
    echo
}

# INFO: checks if input is legal integer
# RETURNS: 0 if int 1 otherwise
is_int () {
    if [[ $(sed -nr '/^-?[[:digit:]]+$/ p' <<< $1) ]];then
        echo "0"
    else
        echo "1"
    fi
}

# INFO: checks if input is legal operator (+ - * /)
# RETURNS: 0 if does 1 otherwise
is_operator () {
    if [[ $(sed -nr '/^[asmd]{1}$/ p' <<< $1) ]];then
        echo "0"
    else
        echo "1"
    fi
}

# =======================================================
# Reverse Polish Notation integer Calculator with 100 symbol stack and 4 operators
# (Very crude version of excellent `dc'); NONITERACTIVE (i/m lazy) && proper
# use of RPN is assumed!!! (just like in excellent `dc')
E_STACKOVERFLOW=49
E_STACKNOTEMPTY=50
E_SYMBOL=51
E_DIVZERO=52
left_operand=
right_operand=
result=

# don't be shmuck, use a stack
for SYMBOL in $*;do
#    echo "SP: $SP, BP: $BP, stack: ${stack[@]}"                        # DEBUG
    if [[ $(is_int $SYMBOL) = "0" ]];then            # put operand on the stack
        push $SYMBOL
        if [[ $? -ne 0 ]];then
            echo "Stack error <$?>."
            exit $?
        fi
    elif [[ $(is_operator $SYMBOL) = "0" ]];then # pop 2 symbols from the stack
        pop
        right_operand="${Data[0]}"
        pop
        left_operand="${Data[0]}"
        case $SYMBOL in
            a) result=$(( left_operand + right_operand));;
            s) result=$(( left_operand - right_operand));;
            m) result=$(( left_operand * right_operand));;
            d)
                if [[ $right_operand -eq 0 ]];then
                    echo "Division by zero error."
                    exit $E_DIVZERO
                else
                    result=$(( left_operand / right_operand ))
                fi
                ;;
        esac
        push $result                                # store result on the stack
        if [[ $? -ne 0 ]];then
            echo "Stack error <$?>."
            exit $?
        fi
    else
        echo "Unknown symbol <$SYMBOL>. Please use RPN symbols only."
        exit $E_SYMBOL
    fi
done

#echo "SP: $SP, BP: $BP, stack: ${stack[@]}"                            # DEBUG
if [[ $(( BP - SP )) -eq 1 ]];then
    pop
    echo "${Data[0]}"  # "$result"
else
    echo "Stack is not empty."
    exit $E_STACKNOTEMPTY
fi

exit 0

# Exercises:
# ---------

# 1)  Modify the "push()" function to permit pushing
#   + multiple element on the stack with a single function call.

# 2)  Modify the "pop()" function to permit popping
#   + multiple element from the stack with a single function call.

# 3)  Add error checking to the critical functions.
#     That is, return an error code, depending on
#   + successful or unsuccessful completion of the operation,
#   + and take appropriate action.

# 4)  Using this script as a starting point,
#   + write a stack-based 4-function calculator.
