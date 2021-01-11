#!/usr/bin/bash

E_NOFILE=55
E_WRONGARG=56
REMOVE_BLANKS=0                                                       # -r flag
BLANKS=0                                         # number of blank lines to add
FILE=
USAGE="Usage: $(basename $0) -f <file-name> [-s <how-much-blank-lines-to-add> | -r]."

while getopts ":f:s:r" OPT;do
    case $OPT in
        f)
            FILE="$OPTARG"
            if [[ ! -f "$FILE" ]];then
                echo "File not found: <$FILE>." 1>&2
                echo "$USAGE" 1>&2
                exit $E_NOFILE
            fi
            ;;
        s)
            BLANKS=$OPTARG;;
        r)
            REMOVE_BLANKS=1;;
        :)
            echo "Unknown option argument: <$OPTARG>." 1>&2
            echo "$USAGE" 1>&2
            exit $E_WRONGARG
            ;;
        \?)
            echo "$USAGE" 1>&2
            exit $E_WRONGARG
            ;;
    esac
done

shift $(( OPTIND - 1 ))

if [[ $REMOVE_BLANKS -ne 0 && $BLANKS -ne 0 ]];then
    echo "Options <-s> and <-r> are mutually exclusive." 1>&2
    echo "$USAGE" 1>&2
    exit $E_WRONGARG
fi

echo;echo

[[ $REMOVE_BLANKS -ne 0 ]] && sed -nr '/^$/d; /^.+/p' "$FILE"

if [[ $BLANKS -ne 0 ]];then
    while read line; do
        echo $line
        for ((i=0; i<BLANKS; i++)){
            echo
        }
    done < "$FILE"
fi

exit $?
