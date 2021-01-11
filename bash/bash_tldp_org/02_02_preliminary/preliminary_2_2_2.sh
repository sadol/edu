#!/bin/bash
# put date users & uptime in stdout & file

LOG_FILE="2_2_2.log"
COMMAND="tee -a ${LOG_FILE}"

# [ -f "$LOG_FILE" ] && cat /dev/null > "$LOG_FILE"  # clear file if exists
#[[ -f "$LOG_FILE" ]] && : > "$LOG_FILE"  # clear file if exists
[[ -f "$LOG_FILE" ]] && > "$LOG_FILE"  # clear file if exists
# touch "$LOG_FILE"  # create it if it does not
date | ${COMMAND}
users | ${COMMAND}
uptime | ${COMMAND}

exit
