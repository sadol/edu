#!/usr/bin/bash

USAGE="$(basename $0) [-l LIMIT] [-q]."
E_ARG=67
E_ROOT=68
E_MAIL=69
ROOT=0
LIMIT=500 #MB
QUOTA=0

if [[ $(id -u) -ne $ROOT ]];then
    echo "You need to be \`root' to run this script."
    echo "$USAGE" >&2
    exit $E_ROOT
fi

while getopts ":l:q:" OPT; do
    case $OPT in
        l)  LIMIT="$OPTARG";;
        q)  QUOTA=1;;
        :)
            echo "Option without argument." >&2
            echo "$USAGE" >&2
            exit $E_ARG
            ;;
        \?)
            echo "Unknown option <$OPT>." >&2
            echo "$USAGE" >&2
            exit $E_ARG
            ;;
    esac
done

shift $(( OPTIND -1 ))

USERS=( $(cat /etc/passwd | cut -f1,3,6 -d: | grep home) )
SUBJECT="Quota warning."
BODY="Default disk quota of <$LIMIT> exeeded."

for user in ${USERS[@]}; do
    data=( ${user//:/ } )
    name="${data[0]}"
    uid="${data[1]}"
    home="${data[2]}"
    mail=mock_function $name                        # should be created somehow
    size="$(df --output='used' -BM -h|grep M)"; size="${size:0:-1}"
    if [[ $size -gt $LIMIT ]];then
        mailx -v -s "$SUBJECT" "$mail" <<< "$BODY"
    fi
done

E_QUOTA=70
SUCCESS=0

# QUOTA jazz
if [[ $QUOTA -eq 1 ]];then
    which quota > /dev/null 2>&1
    if [[ $? -ne $SUCCESS ]];then
        echo "<quota> tools not installed on this system."
        exit $E_QUOTA
    fi
    :  # and so on and on with quota thing
fi

exit $?
