#!/bin/bash
# isalpha.sh: Using a "case" structure to filter a string.

SUCCESS=0
FAILURE=1
shopt -s extglob

isalpha_first () {  # Tests whether *first character* of input string is alphabetic.
    [[ $# -eq 1 ]] || return $FAILURE                # No argument passed?

    case "$1" in
        [[:alpha:]]*) return $SUCCESS;;  # Begins with a letter?
        *           ) return $FAILURE;;
    esac
}             # Compare this with "isalpha ()" function in C.


isalpha_all () {   # Tests whether *entire string* is alphabetic.
    [[ $# -eq 1 ]] || return $FAILURE

    case $1 in
        *[![:alpha:]]*|"") return $FAILURE;;
                        *) return $SUCCESS;;
    esac
}

isdigit ()    # Tests whether *entire string* is numerical.
{             # In other words, tests for integer variable.
    [[ $# -eq 1 ]] || return $FAILURE
    case $1 in
        *[![:digit:]]*|"") return $FAILURE;;
                        *) return $SUCCESS;;
    esac
}

isfloat () {
    [[ $# -eq 1 ]] || return $FAILURE
    case $1 in
                      @(.)+([[:digit:]]) ) return $SUCCESS;;
                      +([[:digit:]])@(.) ) return $SUCCESS;;
        +([[:digit:]])@(.)+([[:digit:]]) ) return $SUCCESS;;
                                        *) return $FAILURE;;
    esac
}

check_var () { # Front-end to isalpha ().
    if isalpha_first "$@";then
        echo "\"$*\" begins with an alpha character."
        if isalpha_all "$@";then
            echo "\"$*\" contains only alpha characters."
        else
            echo "\"$*\" contains at least one non-alpha character."
        fi
    else
        echo "\"$*\" begins with a non-alpha character."
    fi

    echo
}

digit_check () {  # Front-end to isdigit ().
    if isdigit "$@";then
        echo "\"$*\" contains only digits [0 - 9]."
    else
        echo "\"$*\" has at least one non-digit character."
    fi

    echo
}

float_check () {  # Front-end to isfloat ().
    if isfloat "$@";then
        echo "\"$*\" is proper float value (in C sense)."
    else
        echo "\"$*\" isn't proper float value (in C sense)."
    fi

    echo
}

a=23skidoo
b=H3llo
c=-What?
d=What?
e=$(echo $b)   # Command substitution.
f=AbcDef
g=27234
h=27a34
i=27.34
j=.004
k=100.
l=.1.1
m=..3
n=4..
o=4,

check_var $a
check_var $b
check_var $c
check_var $d
check_var $e
check_var $f
check_var     # No argument passed, so what happens?
#
digit_check $g
digit_check $h
digit_check $i
float_check $i
float_check $j
float_check $k
float_check $l
float_check $m
float_check $n
float_check $o

exit 0        # Script improved by S.C.

# Exercise:
# --------
#  Write an 'isfloat ()' function that tests for floating point numbers.
#  Hint: The function duplicates 'isdigit ()',
#+ but adds a test for a mandatory decimal point.
