#!/usr/bin/bash

# to find broken links on a html page use `wget'
SITE=$1
E_NO_SITE=45
if [[ -z $SITE ]]; then
    echo "<SITE> argument is empty." >&2
    exit $E_NO_SITE
fi

wget -r --spider "$SITE"
