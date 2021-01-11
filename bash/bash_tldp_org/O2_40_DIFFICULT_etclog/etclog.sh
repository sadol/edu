#!/usr/bin/bash

# To log every `/etc' rwx you should use `audit' framework in linux, do not
# try to reinvent the wheel.
# With `systemd' distro it is possible to:
#   1. install `audit' package (if not installed yet)
#   2. start (or enable) systemd service: `sudo systemctl start auditd.service'
#   3. add some rules:
#      auditctl -W /etc                       ---> to remove old `/etc' watchers
#      auditctl -w /etc -p rwxa -k etc_access ---> to add new watcher

SUCCESS=0
FAILURE=1
ROOT=0                                                            # root's uid
AUDIT_SERVICE="auditd.service"

E_NO_SYSTEMD=47
which systemctl >& /dev/null
[[ $? -ne $SUCCESS ]] && exit $E_NO_SYSTEMD

# INFO: derem this section if you do not have historical data at hand
# E_NO_SERVICE=46
# systemctl status $AUDIT_SERVICE >& /dev/null
# [[ $? -ne $SUCCESS ]] && exit $E_NO_SERVICE

E_NOT_ROOT=45
[[ "$(id -u)" -ne 0 ]] && exit $E_NOT_ROOT      # only root can run this script

DATAFILE="etc_access.csv"
[[ -f $DATAFILE ]] && rm -rf $DATAFILE
# INFO: START & END must be properly set !!!
ausearch --format csv --start yesterday --end recent -k etc_access > $DATAFILE 2> /dev/null

################################################################################
# INFO:The whole script should end HERE, everything below is slow & unnecesary.#
#      `DATAFILE' should be loaded into a RDBMS of some kind (MariaDB or such) #
#      for further analysis. There is no need to stare at 50k pretty printed   #
#      rows of data.                                                           #
################################################################################

OUTPUTFILE="output.txt"
[[ -f $OUTPUTFILE ]] && rm -rf $OUTPUTFILE
FORMAT=( "TIME" "USER" "TYPE" "FILE" "COMM" )     # output headers (for pprint)
IN_FORMAT=( "TIME" "SUBJ_SEC" "ACTION" "OBJ_PRIME" "HOW" )  # input headers (from csv)
declare -A maxlen # maximum lengths of the respective fields i the output table
declare -A blanks # number of spaces in the field name

for fieldname in ${IN_FORMAT[@]};do
    blanks[$fieldname]=2      # starting point: one space from the right & left
    maxlen[$fieldname]=$(( ${#fieldname} + ${blanks[$fieldname]} ))
done

declare -a ORIG_FORMAT                          # full original header from csv
read line < $DATAFILE                             # load header into a variable
ORIG_FORMAT=( ${line//,/ } )                         # create array from header

# determine widths of the fields
while IFS="," read ${ORIG_FORMAT[@]}; do   # ad-hoc variables with names as in csv data header
    for fieldname in ${IN_FORMAT[@]};do    # using only SOME of the ad-hoc variables defined in bulk above
        fieldvalue=${!fieldname}      # indirect reference of the value of the ad-hoc variable
        [[ $(( ${#fieldvalue} + 2 )) -gt ${maxlen[$fieldname]} ]] && maxlen[$fieldname]=$(( ${#fieldvalue} + 2 ))
    done
done < $DATAFILE

# determine number of spaces in each header
for fieldname in ${IN_FORMAT[@]};do               # using said ad-hoc variables
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
for fieldname in ${IN_FORMAT[@]}; do
    no_leading_spaces=$(( ${blanks[$fieldname]} / 2 ))
    no_trailing_spaces=$(( ${blanks[$fieldname]} - $no_leading_spaces ))
    HEADER+=$(printf "%${no_leading_spaces}s%s%${no_trailing_spaces}s|" " " "$fieldname" " ")
done
HEADER="|${HEADER}"

# 4. build and print data
echo "$vertical" > $OUTPUTFILE
echo "$HEADER" >> $OUTPUTFILE
echo "$vertical" >> $OUTPUTFILE

{                                                   # to remove original header
    read                                            # to remove original header
    while IFS="," read ${ORIG_FORMAT[@]}; do   # ad-hoc variables with names derived from the array `FORMAT' values
        line="|"                             # data line to be contructed
        for fieldname in ${IN_FORMAT[@]};do  # using said ad-hoc variables
            fieldvalue=${!fieldname}         # indirect reference of the value of the ad-hoc variable
            line="${line}""$(printf "%-${maxlen[$fieldname]}s|" "${!fieldname}")"
        done
        printf "$line\n" >> $OUTPUTFILE
    done
} < $DATAFILE                                       # to remove original header

echo "$vertical" >> $OUTPUTFILE

# you can remove original csv here if you wish
# rm -rf $DATAFILE >& /dev/null
echo "Output file \`$OUTPUTFILE' generated."
