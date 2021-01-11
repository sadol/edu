#!/usr/bin/bash

# checks strength of the user`s password; script returns 0 (success) || 1 (faulure)
# INFO: simple english dictionary for passphrase-strength check is not enought,
#       you should equip yourself with proper `dictionary-attack' dictionary
#       full of old and weak passwords as well.
# WARNING: it can be SLOOOOOOOOOOOOOOOOOOOOOOOOOW, so use a random passwords!!!
E_WRONGARG=57
E_NO_DICT=58
base="$(basename $0)"
USGEMSG="Usage: $base [-D <english-dictionary-list>] [-P <pass-dictionary-list>] <passphrase-to-check>;
       for example: $base -D dict1,dict2,dict3 -P dict4,dict5 ."
PASS=""
EN_DICTIONARIES=( "en_dict.txt" ) # as many dicts as you wish (firsts are defaults)
PASS_DICTIONARIES=( "pass_dict.txt" )
EN_DICTIONARY_LIST=""
PASS_DICTIONARY_LIST=""
SUCCESS=0
FAILURE=1
RESULT=0
MIN_LEN=8                                      # minimum length of a passphrase
NO_NUMERIC=1                       # minimal number of numerals in a passphrase
NO_OTHERS=1     # minimal number of non alphanumeric characters in a passphrase

if [[ ! ( -f "${EN_DICTIONARIES[0]}" && -f "${PASS_DICTIONARIES[0]}" ) ]];then
    echo "No default dictionaries detected!" >&2
    exit $E_NO_DICT
fi

while getopts ":d:p:" OPT; do
    case $OPT in
        d)  EN_DICTIONARY_LIST="$OPTARG";;              # set new english dictionary
        p)  PASS_DICTIONARY_LIST="$OPTARG";; # set new known-weak-passwords dictionary
        :)
            echo "Missing argument for option <$OPTARG>." >&2
            echo "$USAGEMSG" >&2
            exit $E_WRONGARG
            ;;
        \?)
            echo "$USAGEMSG" >&2
            exit $E_WRONGARG
            ;;

    esac
done

shift $(( OPTIND - 1 ))
PASS="$1"                                                 # passphrase to check

# initial check of a passphrase
# OUTPUT: <SUCCESS|FAILURE>
_init_check() {
    local ret=$SUCCESS
    local len=${#PASS}
    local temp=
    local temp_len

    # some initial asumptions
    if (( len < MIN_LEN ));then
        ret=$FAILURE
    else
        temp="${PASS//[[:digit:]]/}"
        temp_len=${#temp}
        if (( temp_len < ( MIN_LEN - NO_NUMERIC ) ));then
            ret=$FAILURE
        else
            temp="${PASS//[[:alnum:]]/}"
            temp_len=${#temp}
            (( temp_len < NO_OTHERS )) && ret=$FAILURE
        fi
    fi

    echo $ret
}


# fills dictionary arrays with new values
# INPUT: $1 -> colon separated list of files
#        $2 -> name of the array
_fill_array() {
    local i
    local temp=( ${1//,/ } )                   # extract file names from a list
    local dict

    for (( i=0; i<${#temp[@]}; i++ )) {
        dict="${temp[i]}"
        if [[ ! -f "$dict" ]];then
            echo "No such dictionary <"$dict"> detected." >&2
            exit $E_NO_DICT
        else
            eval $2[i+1]="$dict"
        fi
    }
}


# checks if password contains unsafe substring
# OUTPUT: <SUCCESS|FAILURE>
_check_pass() {
    local dict_id=0
    local ret=$SUCCESS
    local known_pass
    local dict

    for ((dict_id=0; dict_id<${#PASS_DICTIONARIES[@]}; dict_id++)) {
        dict="${PASS_DICTIONARIES[dict_id]}"
        while read known_pass; do
            # additional assumption (do not freak out over very short substrings)
            if (( ${#known_pass} > 3 ));then
                #  native bash regex handling: regex inside " " is treated as
                #+ ordinary string(which is very fortunate in this particular
                #+ case becausee there may be present regex special chars
                #+ inside a line which must be ignored by regex machinery and
                #+ interpreted as simple strings, which is desired)
                if [[ $PASS =~ "$known_pass" ]];then
                    ret=$FAILURE
                    break 2
                fi
            fi
        done < "$dict"
    }

    echo $ret
}

# its easier to make similar function than same kind of universal interface in
# BASH; this function is the same as above but works with different set of
# dictionaries
_check_en() {
    local dict_id=0
    local ret=$SUCCESS
    local word
    local dict

    for ((dict_id=0; dict_id<${#EN_DICTIONARIES[@]}; dict_id++)) {
        dict="${EN_DICTIONARIES[dict_id]}"
        while read word; do
            # additional assumption (do not freak out over very short substrings)
            if (( ${#word} > 3 ));then
                #  native bash regex handling: regex inside " " is treated as
                #+ ordinary string(which is very fortunate in this particular
                #+ case becausee there may be present regex special chars
                #+ inside a line which must be ignored by regex machinery and
                #+ interpreted as simple strings, which is desired)
                if [[ $PASS =~ "$word" ]];then
                    ret=$FAILURE
                    break 2
                fi
            fi
        done < "$dict"
    }

    echo $ret
}

# `main' function; operates on globals
main() {
    RESULT="$(_init_check)"
    if [[ $RESULT -eq $SUCCESS ]];then
        [[ -n $EN_DICTIONARY_LIST ]] && _fill_array "$EN_DICTIONARY_LIST" EN_DICTIONARIES
        [[ -n $PASS_DICTIONARY_LIST ]] && _fill_array "$PASS_DICTIONARY_LIST" PASS_DICTIONARIES
        RESULT="$(_check_pass)"
        [[ $RESULT -eq $SUCCESS ]] && RESULT="$(_check_en)"
    fi
    exit $RESULT
}

#---------------------------------MAIN-----------------------------------------
main
