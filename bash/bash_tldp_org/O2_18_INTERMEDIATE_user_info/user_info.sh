#!/usr/bin/bash

E_WHO=56
E_LASTLOG=57

which who > /dev/null 2>&1
[[ $? -ne 0 ]] && exit $E_WHO
which lastlog> /dev/null 2>&1
[[ $? -ne 0 ]] && exit $E_LASTLOG

LOGGEDUSERS=( $(who | cut -f1 -d" ") )
USERS=( $(cat /etc/passwd | cut -f1,5,6 -d: | grep home) )

for loggeduser in ${LOGGEDUSERS[@]}; do
    for user in ${USERS[@]}; do
        data=( ${user//:/ } )
        login="${data[0]}"
        name="${data[1]}"                                        # may be empty
        [[ -z $name ]] && name="$login"

        if [[ $loggeduser = $login]];then
            info=( $(lastlog -u "$login") )  # tabulaization to remove unnecesary spaces
            date=""
            # weird extraction process; the last 6 elements consist full date
            for ((i=$(( ${#a[@]} - 6 )); i<${#a[@]}; i++ )){
                date="$date ${a[i]}"
            }
            echo "Name(or login): $name, last login: $date"
        fi
    done
done

exit $?
