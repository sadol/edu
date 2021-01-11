#!/usr/bin/bash

# With systemd you should use `loginctl list-users' or such
# (present logins).
# Nicely formatted (present & historical) logins should be obtained by :
MYTIME="now"  # may be `today', `yesterday' etc.
last --present $MYTIME
