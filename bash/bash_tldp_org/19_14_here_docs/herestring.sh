#!/bin/bash
# This script will not run under Bash versions -lt 3.0.

# sed scripts; need `sed -rn' options
sed_new_mail='/^From .+ [[:digit:]]{4}$/ p'
sed_is_header='/^[[:alpha:]-]+: / p'
sed_extract_name='s/^([[:alpha:]-]+): (.+)$/\L\1/p'
sed_extract_value='s/^([[:alpha:]-]+): (.+)$/\2/p'
E_MISSING_ARG=87
E_UNKNOWN_FIELD=88
SUCCESS=0
FAILURE=1

if [[ -z "$1" ]];then
    echo "Usage: $0 mailbox-file"
    exit $E_MISSING_ARG
fi

# DESCRPTION: checks if mail line is a header
# USAGE: is_header <line-of-text>;returns 0 if line is indeed header,
#        1 otherwise
is_header () {
    if [[ $# -eq 1 ]];then
        if [[ $(sed -rn "$sed_is_header" <<< "$1") ]];then
            echo $SUCCESS
        else
            echo $FAILURE
        fi
    fi
}

# DESCRIPTION : checks if there is the new mail indicator in a text file
# USAGE: is_new_mail <line-of-text>; returns 0 if it's new mail,
#        1 otherwise
is_new_mail () {
    if [[ $# -eq 1 ]];then
        if [[ $(sed -rn "$sed_new_mail" <<< "$1") ]];then
            echo $SUCCESS
        else
            echo $FAILURE
        fi
    fi
}

process_mail() {
    local -i new_mail=0
    local header value name date sender body

    while IFS= read -r line;do
#         ^^^^                 Reset $IFS to block `read' from stripping leading and trailing spaces
#                   ^^         do not let `read' to interpret \ character
        # check if it's new email
        if [[ $(is_new_mail "$line") -eq $SUCCESS ]];then            # new mail
            (( new_mail++ ))
        elif [[ $(is_header "$line") -eq $SUCCESS ]];then       # header fields
            name=$(sed -nr "$sed_extract_name" <<< "$line")
            value=$(sed -nr "$sed_extract_value" <<< "$line")
            case "$name" in
                "date")
                    date="$value"
                    ;;
                "from")
                    sender="$value"
                    ;;
                # fill this up if you want more info in the final report
                #-------------------------------------------------------
                "subject")
                    ;;
                "to")
                    ;;
                "message-id")
                    ;;
                "user-agent")
                    ;;
                #-------------------------------------------------------
                *)
                    echo "Unknown mail header <$name>. \
                          Please update script <$(basename $0)>." 1>&2
                    exit $E_UNKNOWN_FIELD
                    ;;
            esac
        else                                                             # body
            #print some info ,if you want to fill `body' variable do it here
            if [[ $new_mail -eq 1 ]];then
                echo "MESSAGE of $date."
                echo "E-Mail address of sender: $sender"
                (( new_mail-- ))
            fi
        fi
    done < "$1"
}

process_mail "$1"  # Send mailbox file to function.

exit $?

# Exercises:
# ---------
# 1) Break the single function, above, into multiple functions,
#+   for the sake of readability.
# 2) Add additional parsing to the script, checking for various keywords.
