#!/usr/bin/bash
byte1=''
byte2=''
byte3=''
byte4=''

for (( ; ; )) {
    echo -n "Please hit a key: "
    read -sn1 byte1
    read -sn1 byte2
    read -sn1 byte3
    case "$byte3" in
        A)  echo "up";;
        B)  echo "down";;
        C)  echo "right";;
        D)  echo "left";;
        *)  # there is one byte more
            read -sn1 byte4
            case "$byte3" in
                1) echo "home";;
                2) echo "insert";;
                3) echo "delete";;
                4) echo "end";;
                5) echo "pgup";;
                6) echo "pgdown";;
            esac
    esac
}
# ========================================= #

#  Exercise:
#  --------
#  1) Add detection of the "Home," "End," "PgUp," and "PgDn" keys
