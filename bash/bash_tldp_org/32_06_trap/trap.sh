#!/bin/bash
# logon.sh: A quick 'n dirty script to check whether you are on-line yet.
# INFO : in modern linux system there is no need to analyse network connection
# with a temp file; `systemd' makes possible to use `systemctl' tool to check
# state of the network service

E_LOG=45
SERVICEFILE="wpa_supplicant-nl80211@wlp5s0.service"   # may vary in your system
SERVICE="${1:-${SERVICEFILE}}"
SERVICEPATH="/etc/systemd/system/${SERVICE}"
if [[ ! -f "$SERVICEPATH" ]];then
    echo "There is no such file <$1> in the filesystem."
    exit $E_LOG
fi

KEYWORD_old="Active: active"
KEYWORD="Active: inactive"
ONLINE=22
OFFLINE=23
USER_INTERRUPT=13

#  Cleans up the temp file if script interrupted by control-c.
trap 'echo I`m cleaning right now!; exit $USER_INTERRUPT' TERM INT

echo

while true ;do
    search="$( grep $KEYWORD <($systemctl status ${SERVICEFILE}) )"

    if [[ ! -z "$search" ]];then #  Quotes necessary because of possible spaces.
        echo "Off-line"
        exit $OFFLINE
    else
        echo -n "."        #  The -n option to echo suppresses newline,
                           #+ so you get continuous rows of dots.
    fi

    sleep 1                                                              # ugly
done


#  Note: if you change the KEYWORD variable to "Exit",
#+ this script can be used while on-line
#+ to check for an unexpected logoff.

# Exercise: Change the script, per the above note,
#           and prettify it.

exit 0


# Nick Drage suggests an alternate method:
# INFO: `ifconfig <device>' is deprecated; use `ip link show <device>' instead
while ! ip link show ppp0 | grep "state UP"; do
    # ifconfig ppp0 | grep UP 1> /dev/null && echo "connected" && exit 0
    echo -n "."   # Prints dots (.....) until connected.
    sleep 2
done
echo "connected"
# Problem: Hitting Control-C to terminate this process may be insufficient.
#+         (Dots may keep on echoing.)
# Exercise: Fix this.



# Stephane Chazelas has yet another alternative:

CHECK_INTERVAL=1

while ! tail -n 1 "$LOGFILE" | grep -q "$KEYWORD"
do echo -n .
   sleep $CHECK_INTERVAL
done
echo "On-line"

# Exercise: Discuss the relative strengths and weaknesses
#           of each of these various approaches.
