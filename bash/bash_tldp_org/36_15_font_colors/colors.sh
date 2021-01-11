#!/usr/bin/bash

foreground=$(tput setaf $1)
background=$(tput setab $2)
if [[ $3 -eq 1 ]];then
    bold=$(tput bold)
else
    bold=""
fi
message="$4"
reset=$(tput sgr0)
echo -n ${foreground}${background}${bold}
printf '%s' "${message}${reset}"
echo
