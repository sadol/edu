#!/bin/bash
# date-calc.sh
# Author: Nathan Coulter
# Used in ABS Guide with permission (thanks!).

MPHR=60    # Minutes per hour.
HPD=24     # Hours per day.
ERRORMSG="usage: $(basename $0) <current-date> <target-date>, \
where date is in format YYYY-MM-DD HH:MM:SS"
E_NUMARGS=95
E_WRONGARG=96
ARGS=2
# VERY crude date-time check (because it's locale dependent)
#REGEX="^[1-2][019][0-9]{2}-[01][0-9]-[012][0-9] [012][0-9]:[0-5][0-9]:[0-5][0-9]$"

if [[ $# -ne $ARGS ]]; then
    echo "${ERRORMSG}"
    exit $NUMARGS
fi

# poor date format check
#if [[ ! $1 =~ $REGEX ]] || [[ ! $2 =~ $REGEX ]];then
#    echo "${ERRORMSG}"
#    exit $E_WRONGARG
#fi


# better format check
date -d"${1}" > /dev/null 2>&1
if [[ $? == 1 ]];then
    echo "First argument is invalid."
    echo "$ERRORMSG"
    exit $E_WRONGARG
fi

date -d"${2}" > /dev/null 2>&1
if [[ $? == 1 ]];then
    echo "Second argument is invalid."
    echo "$ERRORMSG"
    exit $E_WRONGARG
fi

CURRENT=$1
TARGET=$2

diff () {
        printf '%s' $(( $(date -u -d"$TARGET" +%s) -
                        $(date -u -d"$CURRENT" +%s) ))
#                       %d = day of month.
}


#CURRENT=$(date -u -d '2007-09-01 17:30:24' '+%F %T.%N %Z')
#TARGET=$(date -u -d'2007-12-25 12:30:00' '+%F %T.%N %Z')
# %F = full date, %T = %H:%M:%S, %N = nanoseconds, %Z = time zone.

printf '\nIn %s, %s ' \
       "$(date -d"${CURRENT}" '+%Y')" \
       "$(date -d"$CURRENT +
        $(( $(diff) /$MPHR /$MPHR /$HPD / 2 )) days" '+%d %B')"
#       %B = name of month                ^ halfway
printf 'was halfway between %s ' "$(date -d"$CURRENT" '+%d %B')"
printf 'and %s\n' "$(date -d"$TARGET" '+%d %B')"

printf '\nOn %s at %s, there were\n' \
        $(date -u -d"$CURRENT" +%F) $(date -u -d"$CURRENT" +%T)
DAYS=$(( $(diff) / $MPHR / $MPHR / $HPD ))
CURRENT=$(date -d"$CURRENT +$DAYS days" '+%F %T.%N %Z')
HOURS=$(( $(diff) / $MPHR / $MPHR ))
CURRENT=$(date -d"$CURRENT +$HOURS hours" '+%F %T.%N %Z')
MINUTES=$(( $(diff) / $MPHR ))
CURRENT=$(date -d"$CURRENT +$MINUTES minutes" '+%F %T.%N %Z')
printf '%s days, %s hours, ' "$DAYS" "$HOURS"
printf '%s minutes, and %s seconds ' "$MINUTES" "$(diff)"
printf 'until Christmas Dinner!\n\n'

#  Exercise:
#  --------
#  Rewrite the diff () function to accept passed parameters,
#+ rather than using global variables.
