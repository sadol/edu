#!/bin/bash

# Generates a log file in current directory
# from the tail end of /var/log/messages.

# Note: /var/log/messages must be world readable
# if this script invoked by an ordinary user.
#         #root chmod 644 /var/log/messages

MINUTES=20
PERIOD=$((MINUTES*60))
LINES=5
INSTALLLOG="/var/log/pacman.log"    # in my case this log is much more interesting
LOGFILE="logfile"
CMD3="tail -n $LINES $INSTALLLOG | xargs | fmt -s >> $LOGFILE"

watch -n $PERIOD ${CMD3} >> $LOGFILE

exit $?
