#!/usr/bin/bash

# it's a `cron' job script (noninteractive by definition) run every 48 H

E_TRASH=56
TRASH="${HOME}/TRASH"
[[ ! -d "$TRASH" ]] && exit E_TRASH

rm -rf "${TRASH}/*"

exit $?
