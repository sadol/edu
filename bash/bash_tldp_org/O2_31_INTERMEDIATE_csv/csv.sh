#!/usr/bin/bash

# CSV pretty-printing: IT IS COSTLY (algo is not very efficient).
# INFO: There is NO csv file ERROR checking!!!
DATAFILE=""
E_FILE=56
E_WRONGARG=57
E_FILE_MSG="No datafile found: "
USAGEMSG="Usage: $(basename $0) -f <data-file>"
# format should be takenn from another file or argument
FORMAT=( "SURNAME" "NAME" "STREET" "CITY" "STATE" "ZIP_CODE" "TELEPHONE" )
declare -A maxlen # maximum lengths of the respective fields i the output table
declare -A blanks # number of spaces in the field name

for fieldname in ${FORMAT[@]};do
    blanks[$fieldname]=2      # starting point: one space from the right & left
    maxlen[$fieldname]=$(( ${#fieldname} + ${blanks[$fieldname]} ))
done

while getopts ":f:" OPT; do
    case "$OPT" in
        f)  DATAFILE="$OPTARG";;
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

shift $((OPTIND - 1))

if [[ ! -f "$DATAFILE" ]];then
    echo $E_FILE_MSG "$DATAFILE" >&2
    exit $E_FILE
fi

# determine widths of the fields
while IFS="," read ${FORMAT[@]}; do   # ad-hoc variables with names derived from the array `FORMAT' values
    for fieldname in ${FORMAT[@]};do  # using said ad-hoc variables
        fieldvalue=${!fieldname}      # indirect reference of the value of the ad-hoc variable
        [[ $(( ${#fieldvalue} + 2 )) -gt ${maxlen[$fieldname]} ]] && maxlen[$fieldname]=$(( ${#fieldvalue} + 2 ))
    done
done < "$DATAFILE"

# determine number of spaces in each header
for fieldname in ${FORMAT[@]};do  # using said ad-hoc variables
    spaces=$(( ${maxlen[$fieldname]} - ${#fieldname} ))
    [[ $spaces -gt ${blanks[$fieldname]} ]] && blanks[$fieldname]=$spaces
done

#-----------------------creating nicely looking box----------------------------
# 1. calculate table width
box_width=1
for width in ${maxlen[@]}; do
    (( box_width += (width + 1) ))
done

# 2. build vertical line
vertical=""
for (( i=0; i<box_width; i++ )) {
    vertical+="-"
}

# 3. build header with reasonably centered field names
HEADER=""
for fieldname in ${FORMAT[@]}; do
    no_leading_spaces=$(( ${blanks[$fieldname]} / 2 ))
    no_trailing_spaces=$(( ${blanks[$fieldname]} - $no_leading_spaces ))
    HEADER+=$(printf "%${no_leading_spaces}s%s%${no_trailing_spaces}s|" " " "$fieldname" " ")
done
HEADER="|${HEADER}"

# 4. build and print data
echo $vertical
echo "$HEADER"
echo $vertical
while IFS="," read ${FORMAT[@]}; do   # ad-hoc variables with names derived from the array `FORMAT' values
    line="|"                          # data line to be contructed
    for fieldname in ${FORMAT[@]};do  # using said ad-hoc variables
        fieldvalue=${!fieldname}      # indirect reference of the value of the ad-hoc variable
        line="${line}""$(printf "%-${maxlen[$fieldname]}s|" "${!fieldname}")"
    done
    printf "$line\n"
done < "$DATAFILE"
echo $vertical
#------------------------------------------------------------------------------
exit $?
