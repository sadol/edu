#!/bin/bash
# paragraph-space.sh
# Ver. 2.1, Reldate 29Jul12 [fixup]

# Inserts a blank line between paragraphs of a single-spaced text file.
# Usage: $0 <FILENAME
LEN=0
LINE=""
MINLEN=60        # Change this value? It's a judgment call.
THISISTHEEND=0
#  Assume lines shorter than $MINLEN characters ending in a period
#+ terminate a paragraph. See exercises below.

for (( ; ; )) {
    read LINE || THISISTHEEND=1     # of the world as we know it, i'm feel fine
    LEN=${#LINE}
    if [[ THISISTHEEND -ne 1 ]]; then
        [[ $LEN -lt $MINLEN && "$LINE" =~ [*\.\?\"\!]$ ]] && echo
    else
        exit
    fi
}

exit  # just for fun

# Exercises:
# ---------
#  1) The script usually inserts a blank line at the end
#+    of the target file. Fix this.
#  2) Line 17 only considers periods as sentence terminators.
#     Modify this to include other common end-of-sentence characters,
#+    such as ?, !, and ".
