#!/usr/bin/bash

# Rogue user monitoring tool.
# In fact you do not have to use bash at all if you don't want to.
# `systemd' service `audit.service' should do heavy lifting instead.
# Install `audit' package and read `man auditctl'.
E_NO_AUDIT=45
E_NO_SERVICE=46
E_NO_USER=47
E_NO_KEY=48
SUCCESS=0
FAILURE=1
USER=$1                                                 # no need for `getopts'
KEY=$2                   # the easiest way of doing audit searching is by a key

which auditctl >& /dev/null         # check if `audit' is present in the system
if [[ $? -eq $FAILURE ]];then
    echo "No \`auditctl' command found. Please install \`audit' package."
    exit $E_NO_AUDIT
fi

systemctl status auditd.service >& /dev/null    # check if `audit' service is up
if [[ $? -eq $FAILURE ]];then
    echo " \`auditd.service' is not runnig. Please fire up \`auditd.service'."
    exit $E_NO_SERVICE
fi

id -u "$USER" >& /dev/null
if [[ $? -eq $FAILURE ]];then
    echo "User \`$USER' does not exist."
    exit $E_NO_USER
fi

if [[ -z $KEY ]];then
    echo " \`autit' key <$KEY> does not exist."
    exit $E_NO_KEY
fi

# INFO: read `man auditctl' & `man audit.rules' first!
# aureport -u $USER                                       # not implemented yet
ausearch -k $KEY                                                        # VOILA
