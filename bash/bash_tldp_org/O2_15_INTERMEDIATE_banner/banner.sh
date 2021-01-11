#!/usr/bin/bash

# i do not recall what was going on with that banner thing
DEFAULTBANNER="$(uname -a)"
DEFAULTBANNER="$DEFAULTBANNER

     X       X
    X X      X
   X   X     X
  XXXXXXX    X
 X       X   X
X         X  X
"

BANNERFILE="${1}"
if [[ -f "$BANNERFILE" ]];then
    BANNER="$( cat "$BANNERFILE" )"
else
    BANNER="$DEFAULTBANNER"
fi

echo "$BANNER"
