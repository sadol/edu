#!/usr/bin/bash

# try to adapt this as a `cron' job or `systemd' timer & service
SUCCESS=0
E_NO_SQLITE3=45
E_NO_SENDMAIL=46
E_NO_CONFIG_SENDMAIL=47
PS_DATA="ps.output"
SQL_DATABASE="ps.db"
SQL_OUTPUT="children.out"
TABLE_NAME="processes"
CHILDREN=5
MAIL_RECIPIENT="$(whoami)"        # change this accordingly | use as script arg
SQL="CREATE TABLE $TABLE_NAME
    (pid integer PRIMARY KEY, ppid integer, command varchar);
.separator ","
.import $PS_DATA $TABLE_NAME
.headers on
.mode table
.output $SQL_OUTPUT
SELECT p2.pid as pid, p2.command as process, count(p1.ppid) as children
    from processes p1 INNER JOIN processes p2 ON p1.ppid=p2.pid GROUP BY p1.ppid
    HAVING children>${CHILDREN} ORDER BY children DESC;
.exit
"

# pre-cleaning
[[ -f $PS_DATA ]] && rm -rf $PS_DATA
[[ -f $SQL_DATABASE ]] && rm -rf $SQL_DATABASE
[[ -f $SQL_OUTPUT ]] && rm -rf $SQL_OUTPUT

# simple check
which sqlite3 >& /dev/null
[[ $? -ne $SUCCESS ]] && exit $E_NO_SQLITE3                          # silently
which sendmail >& /dev/null
[[ $? -ne $SUCCESS ]] && exit $E_NO_SENDMAIL                         # silently

# make some data
ps --no-headers axo pid,ppid,comm | sed -nr 's/([ ]+)([^ ]+)([ ]+)([^ ]+)([ ]+)(.*)/\2,\4,\6/p' > $PS_DATA
# load data into database, make query and spit out output
sqlite3 --batch $SQL_DATABASE <<< "$SQL"

# post-cleaning
rm -rf $PS_DATA
rm -rf $SQL_DATABASE

# do smth with output datafile
SUBJECT="Too_much_forking."
BODY="

Processes with more than $CHILDREN children in $(date) :
$(<SQL_OUTPUT)

"

mail -s "$SUBJECT" "$RECIPIENT" <<< "$BODY" 2> /dev/null
if [[ $? -eq $SUCCESS ]];then
    rm -rf "$SQL_OUTPUT"
else
    # clear `SQL_OUTPUT' by yourself in case of mail error
    exit $E_NO_CONFIG_SENDMAIL
fi

exit $?
