#!/usr/bin/bash

# `zenity' in use (bash for lamers;))

# checks package availibility
check_package () {
    if [[ -n "$(which $1| sed -nr '/-bash:/ p')" ]];then
        echo "`$1` package not installed." 2>&1
        exit $2
    fi
}

E_ZENITY=88
E_XRANDR=90
check_package "zenity" "$E_ZENITY"
check_package "xrandr" "$E_XRANDR"

resolution="$(xrandr | sed -nr \
    's/.+current ([[:digit:]]+) x ([[:digit:]]+).+/\1 \2/p')"
HEIGHT="$(( $(cut -d" " -f2 <<< "$resolution") / 2 ))"
WIDTH="$(( $(cut -d" " -f1 <<< "$resolution") / 2 ))"
zenity --text-info --title="$(basename $0)" --filename="$0" --width="$WIDTH" \
    --height="$HEIGHT"

OUTFILE="$0.output"
E_INPUT=89
echo -n "variable=" > "$OUTFILE"
zenity --entry --title="User input" --entry-text="Enter variable please." \
    --text="Value:" --height="$HEIGHT" --width="$WIDTH" >> "$OUTFILE"
if [[ $? -eq 0 ]];then
    echo "<zenity entry> widget executed propery."
else
    echo "Errors, errors everywhere!"
    rm -f "$OUTFILE"
    exit $E_INPUT
fi

# retrieve variable from temp file:
. "$OUTFILE"
echo "Variable value entered: $variable"
rm -f "$OUTFILE"
exit $?
