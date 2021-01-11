#!/usr/bin/bash

# Most fonts from the Unicode table are not properly printed at least at my
# terminals making this exercise completely unappealing.
FILE="UTF.txt"
[[ -f "$FILE" ]] && rm -rf "$FILE"

exec 3>&1
exec >${FILE}

MAXNUM=10000
COLUMNS=5
LITTLESPACE=-7
BIGSPACE=-7

for ((i=1; i<MAXNUM; i++)) {
    paddi="       $i"
    echo -n "${paddi: $BIGSPACE}  "
    paddu="$(echo -e "       \u$(printf '%x' $i)  ")"
    echo -ne "${paddu: $LITTLESPACE}"
    echo -n "    "
    (( i % COLUMNS == 0 )) && echo
}

exec >&3; exec 3>&-
echo; echo "File <$FILE> created."; echo

exit $?
