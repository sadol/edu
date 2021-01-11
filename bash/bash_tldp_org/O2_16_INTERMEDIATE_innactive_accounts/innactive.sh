#!/usr/bin/bash

# this script shoud not be intercative to be used as a cron job
E_ROOT=56
[[ $(id -u) -ne 0 ]] && exit $E_ROOT

# only ids>1000
MINUID=1000                                        #  minimal effective user id
MAXUID=4000                                        # magic number
INACTIVE=90
USERS=( $(lastlog -u $MINUID-$MAXUID -b $INACTIVE | cut -f1 -d" ") )

for (( i=1; i<${#USERS[@]}; i++)) {                      # i=1 to ignore header
    userdel "${USERS[i]}" > dev/null 2>&1 # main loop can be expanded accordingly
}

exit $?
