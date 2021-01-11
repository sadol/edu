#!/usr/bin/bash

# nonintercative script for use by cron and such
EXCLUDE="${HOME}/.cache/*"                                 # rando browser shit
NEWERTHAN=$(( 24 * 60 ))                                   # only the last 24 h
TARBALL="${HOME}_$(date -I)"
# sedddddddddddddd tiiiiiiiiiiiiiiiiiime
TARBALL=$(sed -nr 's/[/-]/_/g; s/^_//p' <<< "$TARBALL")

# use `xargs' to properly deal with wierd chars in the names of files
find "$HOME" -type f -not -path "$EXCLUDE" -print0 -mmin -$NEWERTHAN | \
    xargs -0 tar czf "$TARBALL"
