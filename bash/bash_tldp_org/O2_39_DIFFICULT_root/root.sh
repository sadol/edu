#!/usr/bin/bash

E_NO_BC_SCRIPT=45
E_ARG=46
BC_FILE="root.bc"                    # `bc' script file filled with definitions
BC_MSG="No \`bc' script file <$BC_FILE> found."
USAGE="USAGE: $(basename $0) <positive-number>."
ARG="${1//[[:space:]]/}"                                     # no spaces needed

if [[ ! -f $BC_FILE ]];then
    echo "$BC_MSG" >&2
    echo "$USAGE" >&2
    exit $E_NO_BC_SCRIPT
fi

if [[ ${ARG:0:1} == "-" ]];then  # no negatives needed (for this time at least)
    echo "$USAGE" >&2
    exit $E_ARG
fi

if [[ -z "${ARG}" ]];then
    echo "$USAGE" >&2
    exit $E_ARG
fi

root="$(bc $BC_FILE <<< "scale=7; root(${ARG})")"

echo "$root"
