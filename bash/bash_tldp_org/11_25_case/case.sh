#!/usr/bin/bash

# reading chars from keyboard
EXIT="X"
char=""

for (( ; ; )) {
    echo -n "Type sth ( $EXIT to exit ):"
    read line
    for (( i=0; i<${#line}; i++ )) {
        char=${line:${i}:1}
        case "$char" in
            [[:lower:]] ) echo "Lower case";;
            [[:upper:]] )
                if [[ $char = $EXIT ]]; then
                    echo "Bye!!!"
                    break 2
                fi
                echo "Upper case"
                ;;
            [0-9] ) echo "Number" ;;
            * ) echo "Special char" ;;
        esac
    }
    echo "---------------------"
}

exit $?
