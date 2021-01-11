#!/bin/bash
# fetch_address-2.sh
# A more elaborate version of fetch_address.sh.

SUCCESS=0
E_DB=99    # Error code for missing entry.

declare -A address
#       -A option declares associative array.
usage="usage : $(basename $0) <db_file_path> <name_to_check>"
input="$1"
name="$2"
[[ ! -f "$input" || -z "$name" ]] && echo "$usage" && exit $E_DB

store_address (){
    while IFS="|" read -r nam addr;do
        address[$nam]="$addr"
    done < "$input"
}

# returns empty line in case of the missing name
fetch_address (){
    [[ -n "${address[$1]}" ]] && echo "$1's address is ${address[$1]}." || echo
}


#  Exercise:
#  Rewrite the above store_address calls to read data from a file,
#+ then assign field 1 to name, field 2 to address in the array.
#  Each line in the file would have a format corresponding to the above.
#  Use a while-read loop to read from file, sed or awk to parse the fields.
store_address
fetch_address "$name"
