#!/bin/bash
# Killing ppp to force a log-off.
# For dialup connection, of course.

# Script should be run as root user.
E_NOTRUNNIG=98
E_NOROOT=99
SERPORT=ttyS3
USER=$(whoami)
[[ $USER != "root" ]] && exit $E_NOROOT
#                     -------- process ID of ppp -------
PPPPID=$(ps ax|awk '/ppp/ { print $1 }'|head -n1)
PPPSTATUS=$(ps ax --no-header -q $PPPPID|awk '{print $3}')
# zombie check
[[ $PPPSTATUS =~ [Z] ]] && exit $E_NOTRUNNING
#  Depending on the hardware and even the kernel version,
#+ the modem port on your machine may be different --
#+ /dev/ttyS1 or /dev/ttyS2.

killppp="eval kill -9 $PPPPID"
$killppp                     # This variable is now a command.

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
