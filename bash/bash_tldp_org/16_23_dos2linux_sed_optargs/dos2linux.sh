#!/bin/bash
# Du.sh: DOS to UNIX text file converter.

E_WRONGARGS=85
DOSSUFFIX="dos"
LINUXSUFFIX="linux"
CR='\015'  # Carriage return.
           # 015 is octal ASCII code for CR.
           # Lines in a DOS text file end in CR-LF.
           # Lines in a UNIX text file end in LF only.
LF='\012'
USAGEMSG="Usage: $(basename $0) [-r] <filename>"
FILENAME=""
REVERSE=0

while getopts ":r" OPT ;do
    case "$OPT" in
        r)  # reverse - from linux to dos conversion
            REVERSE=1
            shift
            ;;
        :)
            echo "Unknown option <$OPTARG>." 1>&2
            exit $E_WRONGARG
            ;;
        \?)
            echo "$USAGEMSG" 1>&2
            exit $E_WRONGARG
            ;;
    esac
done

if [[ -z "$1" ]]; then
    echo "File name argument missing." 1>&2
    echo "$USAGEMSG" 1>&2
    exit $E_WRONGARGS
else
    FILENAME="$1"
fi

if [[ ! -f "$FILENAME" ]]; then
    echo "File name argument <$FILENAME> not a regular file." 1>&2
    echo "$USAGEMSG" 1>&2
    exit $E_WRONGARGS
fi

case $REVERSE in
    0)
        # Delete CR's and write to new file.
        tr -d $CR < "$FILENAME" > "$FILENAME.$LINUXSUFFIX"
        echo "Original DOS text file is <$FILENAME>."
        echo "Converted LINUX text file is <$FILENAME.$LINUXSUFFIX> ."
        ;;
    1)
        # below the list of failures:
        # sed 's/\012/\015\012/g' "$FILENAME" > "$FILENAME.$DOSSUFFIX"
        # sed 's/\n/\r\n/g' "$FILENAME" > "$FILENAME.$DOSSUFFIX"
        # sed -i.$DOSSUFFIX 's/$/\r/g' "$FILENAME"
        # and one success (i am almost certain that is not portable)
        sed 's/$/\r/g' "$FILENAME" > "$FILENAME.$DOSSUFFIX"
        echo "Original LINUX text file is <$FILENAME>."
        echo "Converted DOS text file is <$FILENAME.$DOSSUFFIX> ."
        ;;
esac

exit $?

# Exercise:
# --------
# Change the above script to convert from UNIX to DOS.
