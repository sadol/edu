#!/bin/bash
# Killing ppp to force a log-off.
# For dialup connection, of course.
# fuser version
# Script should be run as root user.
E_FUSER=96
E_NOMODEM=97
E_NOTRUNNIG=98
E_NOROOT=99
USER=$(whoami)
SERPORT="ttyS3"
MODEMFILE="/dev/modem"
[[ $USER != "root" ]] && exit $E_NOROOT
[[ ! -c $MODEMFILE ]] && exit $E_NOMODEM

fuser -s -k $MODEMFILE

chmod 666 /dev/$SERPORT      # Restore r+w permissions, or else what?
#  Since doing a SIGKILL on ppp changed the permissions on the serial port,
#+ we restore permissions to previous state.

rm /var/lock/LCK..$SERPORT   # Remove the serial port lock file. Why?

exit $?

# Exercises:
# ---------
# 1) Have script check whether root user is invoking it.
# 2) Do a check on whether the process to be killed
#+   is actually running before attempting to kill it.
# 3) Write an alternate version of this script based on 'fuser':
#+      if [ fuser -s /dev/modem ]; then . . .
