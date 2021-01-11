#!/bin/bash
# script-detector.sh: Detects scripts within a directory.

ISSCRIPTCHARS=2    # Test first 2 characters.
ISBASH="bash"
SHABANG='#!'   # Scripts begin with a "sha-bang."
E_NODIR=96
E_INVALIDOPT=97
NOOFARGS=1
DIRTOSEARCH=""
INFOMSG="Usage: $(basename $0) [-d [dir-to-search]] ."

while getopts ":d:" OPT; do
    case ${OPT} in
        d)
            if [[ ! -d "${OPTARG}" ]];then
                echo "<${OPTARG}> is not a directory." 1>&2
                echo $INFOMSG 1>&2
                exit $E_NODIR
            fi

            DIRTOSEARCH="${OPTARG}*"
            ;;
        :)
            DIRTOSEARCH="*"
            ;;
        \?)
            echo "Invalid option <${OPT}>." 1>&2
            echo $INFOMSG 1>&2
            exit $E_INVALIDOPT
    esac
done

[[ $OPTIND -eq 1 ]] && DIRTOSEARCH="*"

shift $((OPTIND - 1))   # just for the sake of good practice

for file in $DIRTOSEARCH;do
    if [[ -f "$file" ]];then
        if [[ $(head -c$ISSCRIPTCHARS "$file") = "$SHABANG" ]];then
            if [[ $(head -1 "$file" | grep $ISBASH) ]];then
                echo "File \"$file\" is a BASH script."
            else
                echo "File \"$file\" is an unknown script."
            fi
        else
            echo "File \"$file\" is *not* a script."
        fi
    fi
done

exit 0

#  Exercises:
#  ---------
#  1) Modify this script to take as an optional argument
#+    the directory to scan for scripts
#+    (rather than just the current working directory).
#
#  2) As it stands, this script gives "false positives" for
#+    Perl, awk, and other scripting language scripts.
#     Correct this.
