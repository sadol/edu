#!/usr/bin/bash

E_BADARGS=67
E_NOSSL=68
E_FILE=69
USAGE="Usage: $(basename $0) < -f <file_name> | -e >.
  \`-e' --> echoes 5 pseudorandom numbers from 1 to 50 on \`stdout';
  \`-f <file_name> --> sends said numbers to the dedicated file.
   Both options are mutually exclusive.
  \`OpenSSL' package must be installed on the system."
SUCCESS=0
FILENAME=""
ECHOFLAG=0
declare -a numbers
MINNUM=1
MAXNUM=51
HOWMANY=5

#-------------OpenSSL test---------------------
which openssl > /dev/null 2>&1

if [[ $? -ne $SUCCESS ]];then
    echo "$USAGE" >&2
    exit $E_NOSSL
fi

if [[ $# -gt 2 || $# -lt 1 ]];then
    echo "$USAGE" >&2
    exit $E_BADARGS
fi

#-------------and the rest----------------------
while getopts ':f:e' OPT;do
    case $OPT in
        f)  FILENAME="$OPTARG";;
        e)  ECHOFLAG=1;;
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

if [[ $ECHOFLAG -eq 1 && -n "$FILENAME" ]];then
    echo "$USAGE" >&2
    exit $E_BADARGS
fi

if [[ -f "$FILENAME" ]];then
    echo "File \`$FILENAME' is present in the filesystem." >$2
    echo "$USAGE" >&2
    exit $E_FILE
fi

# return random number from a given range (no error checking in internal
# function !!!)
__gimi_number() {
    local startn stopn temp flagoff
    startn=$1
    stopn=$2
    flagoff=1

    openssl rand -rand "/dev/random"                                # seed PRNG

    # this loop is lame & should be replaced by some clever onliner
    temp="$(( $(printf '%d' "0x$(openssl rand -hex 4)") % $stopn ))"
    until [[ $temp -ne 0 ]];do
        temp="$(( $(printf '%d' "0x$(openssl rand -hex 4)") % $stopn ))"
    done

    echo "$temp"
}

# check duplicates in the `numbers' array
__check_dups() {
    local unique="$(sort < <(for item in ${numbers[@]};do echo $item; done) | uniq | wc -l)"
    echo "$(( HOWMANY - unique ))"
}

# tries to populate the \`numbers' array with 5 pseudorandom numbers
__try_to_populate_numbers() {
    local i
    for ((i=0; i<$1; i++)) {
        numbers[i]="$(__gimi_number $MINNUM $MAXNUM)"
    }
}

populate_numbers() {
    __try_to_populate_numbers $1
    while (( $(__check_dups) )); do
        __try_to_populate_numbers $1
    done
}

populate_numbers $HOWMANY
if [[ $ECHOFLAG -eq 1 ]];then
    echo ${numbers[@]}
elif [[ -n "$FILENAME" ]];then
    date -I > "$FILENAME"
    echo ${numbers[@]} >> "$FILENAME"
fi
