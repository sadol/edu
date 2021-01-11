#!/usr/bin/bash

TRASH="${HOME}/TRASH"
[[ ! -d $TRASH ]] && mkdir "$TRASH" > /dev/null 2>&1

E_MKDIR=34
[[ $? -ne 0 ]] && exit $E_MKDIR

E_EMPTY=35
[[ $# -eq 0 ]] && exit $EMPTY

for file in $@; do
    [[ ! -e $file ]] && continue
    # dont use `awk', use builtins & grammar instead
    data=( $(file --mime-type "$file") )

    if [[ ! ${data[1]} =~ gzip ]];then
        bzname="${file}_$(date -Is).gzip" # to avoid name conflicts in the trashbin
        tar --remove-files -czf "$bzname" "$file"
    else
        bzname="$(date -Is)_${file}" # to avoid name conflicts in the trashbin
    fi

    mv "$bzname" "$TRASH"
done

exit $?
