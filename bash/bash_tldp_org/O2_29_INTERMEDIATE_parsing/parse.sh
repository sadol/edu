#!/usr/bin/bash

# parsing passwd (and formating which is so boring)
FILE="/etc/passwd"
titles=( "login" "pass" "uid" "gid" "name" "home_directory" "interpreter" )
max_widths=( 0 0 0 0 0 0 0 )                         # max widths of the fields
NO_FIELDS=7

# determinig what is the proper field width (did i say i hate formating?)
while IFS=: read ${titles[@]};do
    login="${login// /_}"
    pass="${pass// /_}"
    uid="${uid// /_}"
    gid="${gid// /_}"
    name="${name// /_}"
    home_directory="${home_directory// /_}"
    interpreter="${interpreter// /_}"
    [[ ${max_widths[0]} -lt ${#login} ]] && max_widths[0]=${#login}
    [[ ${max_widths[1]} -lt ${#pass} ]] && max_widths[1]=${#pass}
    [[ ${max_widths[2]} -lt ${#uid} ]] && max_widths[2]=${#uid}
    [[ ${max_widths[3]} -lt ${#gid} ]] && max_widths[3]=${#gid}
    [[ ${max_widths[4]} -lt ${#name} ]] && max_widths[4]=${#name}
    [[ ${max_widths[5]} -lt ${#home_directory} ]] && max_widths[5]=${#home_directory}
    [[ ${max_widths[6]} -lt ${#interpreter} ]] && max_widths[6]=${#interpreter}
done < "$FILE"

# width's title correction
for ((i=0; i<NO_FIELDS; i++)) {
    [[ ${max_widths[i]} -lt ${#titles[i]} ]] && max_widths[i]=${#titles[i]}
}

box_width=15
for((i=0; i<NO_FIELDS; i++)) {
    (( box_width += ${max_widths[i]} ))
}

box() {
    for ((i=0; i<box_width; i++)) {
        echo -n "-"
    }
    echo
}

box
echo -n "|"
for ((i=0; i<NO_FIELDS; i++)) {
    printf '%*s |' ${max_widths[i]} "${titles[i]}"
}
echo

box

while IFS=: read ${titles[@]};do
    echo -n "|"
    printf '%*s |' ${max_widths[0]} "$login"
    printf '%*s |' ${max_widths[1]} "$pass"
    printf '%*s |' ${max_widths[2]} "$uid"
    printf '%*s |' ${max_widths[3]} "$gid"
    printf '%*s |' ${max_widths[4]} "$name"
    printf '%*s |' ${max_widths[5]} "$home_directory"
    printf '%*s |\n' ${max_widths[6]} "$interpreter"
done < "$FILE"

box
