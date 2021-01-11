#!/bin/bash
# Changes a file to all uppercase.

E_BADARGS=85
ERRORMSG="Usage: $(basename $0) -c [ul] -f filename"
FILENAME=""
CASE=""
BAKSUFFIX=".tr"

while getopts ":c:f:" OPT;do
    case "$OPT" in
        c)
            CASE="$OPTARG"
            ;;
        f)
            FILENAME="$OPTARG"
            ;;
        :)
            echo "Option <$OPTARG> needs an argument." 1>&2
            ;;
        \?)
            echo "$ERRORMSG" 1>&2
            exit $E_BADARGS
            ;;
    esac
done

shift $((OPTIND - 1))

if [[ ! -e "$FILENAME" ]]; then
    echo "There is no such file <$FILENAME> in the filesystem." 1>&2
    echo "$ERRORMSG" 1>&2
    exit $E_BADARGS
fi

case $CASE in
    l)
        tr [:upper:] [:lower:] < "$FILENAME" > "$FILENAME""$BAKSUFFIX"
        ;;
    u)
        tr [:lower:] [:upper:] < "$FILENAME" > "$FILENAME""$BAKSUFFIX"
        ;;
    *)
        echo "$ERRORMSG" 1>&2
        exit $E_BADARGS
        ;;
esac

exit $?

#  Exercise:
#  Rewrite this script to give the option of changing a file
#+ to *either* upper or lowercase.
#  Hint: Use either the "case" or "select" command.
