#!/bin/bash
# unit-conversion.sh
# Must have 'units' utility installed.
# VERY STUPID wrapper around quite good `units' utility
N_OF_VALS=2
USAGE="Usage: $(basename $0) -f <from-units-name> -t <to-units-name> [from-unit-val]
Please read 'units' manual in case of troubles with proper unit names."
E_ARGTYPE=45
E_VALS=46
E_UNITS=47

# not very precise check but what the heck...
check_units_return () {
    if [[ -z $1 || $(sed -rn '/unit/p;/error/p;/_/p;/^$/p' <<< $1 ) ]];then # error
        echo "'units' utility threw error. Check 'units' man to find proper unit names." 1>&2
        exit $E_UNITS
    fi
}

while getopts ":f:t:" OPT;do
    case $OPT in
        f) # from units
            FROM_UNITS="$(sed- nr 's/.+/\L&/p' <<< $OPTARG)" # sed RULES!!!
            ;;
        t) # to units
            TO_UNITS="$(sed- nr 's/.+/\L&/p' <<< $OPTARG)"
            ;;
        :)
            echo "Option <$OPT> requires an argument." 1>&2
            echo $USAGE 1>&2
            exit $E_ARGTYPE
            ;;
        \?)
            echo "Invalid option <$OPT>." 1>&2
            echo $USAGE 1>&2
            exit $E_ARGTYPE
            ;;
    esac
done

shift $((OPTIND - 1))

case $# in
    0) # default; returns only multiplier
        RET=$(units "$FROM_UNITS" "$TO_UNITS" | sed -rn '1 {s/^.+ (.+)$/\1/p}')
        check_units_return $RET
        RET="One $FROM_UNITS is equal to $RET $TO_UNITS."
        ;;
    1) # normal case; returns certain value
        VALUE="$1"
        RET=$(units "$VALUE $FROM_UNITS" $TO_UNITS | sed -rn '1 {s/^.+ (.+)$/\1/p}')
        check_units_return $RET
        RET="$VALUE $FROM_UNITS is equal to $RET $TO_UNITS."
        shift
        ;;
    *) # error; returns whatever returns
        echo "Invalid number of terminal arguments <$#>." 1>&2
        echo $USAGE 1>&2
        exit $E_ARGTYPE
        ;;
esac

echo $RET

exit $?


#  What happens if you pass incompatible units,
#+ such as "acres" and "miles" to the function?

# Exercise: Edit this script to accept command-line parameters,
#           with appropriate error checking, of course.
