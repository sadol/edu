#!/bin/bash
# grp.sh: Rudimentary reimplementation of grep.

E_BADARGS=85
E_NODIR=86
E_BADPATTERN=87
PATTERN=""
DIRECTORY="*"
ERRORMSG="Usage: $(basename $0) -p <pattern> [-d <directory>] ."

while getopts ":d:p:" OPT; do
    case "${OPT}" in
        d)
            if [[ ! -d "${OPTARG}" ]];then
                echo "<${OPTARG}> is not a directory." 1>&2
                echo "$ERRORMSG" 1>&2
                exit $E_NODIR
            fi

            if [[ "$OPTARG" = "." ]];then # is it necessary? idk
                DIRECTORY="*"
            else
                DIRECTORY="${OPTARG}*"
            fi
            ;;

        p)
            PATTERN="${OPTARG}"
            ;;
        :)
            echo "Option without argument." 1>&2
            echo "$ERRORMSG" 1>&2
            exit $E_BADARGS
            ;;
        \?)
            echo "Unknown option <$OPT>." 1>&2
            echo "$ERRORMSG" 1>&2
            exit $E_BADARGS
    esac
done

shift $((OPTIND - 1))  # cargo cult code :)

if [[ -z "$PATTERN" ]];then
    echo "No pattern supplied." 1>&2
    echo "$ERRORMSG" 1>&2
    exit $E_BADPATTERN
fi

for file in $DIRECTORY; do
    output="$(sed -n /"$PATTERN"/p $file)"  # Command substitution.
    if [[ ! -z "$output" ]];then           # What happens if "$output" is not quoted?
        while read line;do
            echo -n "$file: "
            echo "$line"
        done < <(printf "%s\n" "${output}")  # process substitution
    fi              #  sed -ne "/$1/s|^|${file}: |p"  is equivalent to above.
done

exit 0

# Exercises:
# ---------
# 1) Add newlines to output, if more than one match in any given file.
# 2) Add features.
