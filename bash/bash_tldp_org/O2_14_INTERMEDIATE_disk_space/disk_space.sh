#!/usr/bin/bash

E_BADARGS=78
DIRECTORY="${HOME}"                                            # starting point
SIZE=100                                                           # how many K
MINSIZE=1
MAXSIZE=10000                                                    # magic number
LOGFILE="deletions_$(date -I)"
USAGE="$(basename $0) [-d <directory] [-s <size-in-Kb>] [ -f <logname>]."

while getopts "':d:s:f:" OPT; do
    case $OPT in
        d)  DIRECTORY="$OPTARG";;
        s)  SIZE="$OPTARG";;
        f)  LOGFILE="$OPTARG";;
        :)
            echo "Option without argument." >&2
            echo "$USAGE" >&2
            exit $E_BADARGS
            ;;
        \?)
            echo "Unknown option <$OPT>." >&2
            echo "$USAGE" >&2
            exit $E_BADARGS
            ;;
    esac
done

shift $(( OPTIND - 1 ))

E_DIR=79
if [[ ! -d "$DIRECTORY" ]];then
    echo "Unknown directory: <${DIRECTORY}>." >&2
    echo "$USAGE" >&2
    exit $E_DIR
fi

E_SIZE=80
if [[ $SIZE -lt $MINSIZE || $SIZE -gt $MAXSIZE ]];then
    echo "Illegal size: <${SIZE}>." >&2
    echo "$USAGE" >&2
    exit $E_SIZE
fi

E_LOG=81
if [[ -f "$LOGFILE" ]];then
    echo "Logfile is present in the filesystem: <${LOGFILE}>." >&2
    echo "$USAGE" >&2
    exit $E_LOG
fi

>"$LOGFILE"

# lets create array of files to remove or compress to avoid `exec' dance
FILES=( $(find "$DIRECTORY" -type f -size "+${SIZE}k" -print0|xargs -0 ls) )

for file in ${FILES[@]}; do
    rm -i "$file"
    if [[ ! -f "$file" ]];then
        printf '%s DELETED %s\n' "$file" "$(date -Is)" >> "$LOGFILE"
        continue                                            # no tarball needed
    fi

    bzname="${file}.bzip"
    echo -n "Add \`$file\' to the \`$bzname\' archive ? : "
    read ans
    [[ $ans = "y" ]] && tar --remove-files -czf "$bzname" "$file"
done
