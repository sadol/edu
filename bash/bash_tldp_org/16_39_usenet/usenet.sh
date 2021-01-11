#!/bin/bash
# Uudecodes all uuencoded files in current working directory.

lines=35        # Allow 35 lines for the header (very generous).
sedscript='1,'$lines' {
                       /subject/I p
}
'

for File in *;do   # Test all the files in $PWD.
    search1="$(head -n $lines $File | grep begin | wc -w)"
    search2="$(tail -n $lines $File | grep end | wc -w)"
    search3="$(sed -nr "sedscript" <<< "$File"|wc -l)"
    #  Uuencoded files have a "begin" near the beginning,
    #+ and an "end" near the end.
    if [[ "$search1" -gt 0 ]];then
        if [[ "$search3" -gt 0 ]];then
            if [[ "$search2" -gt 0 ]];then
                echo "uudecoding - $File -"
                uudecode $File
            fi
        fi
    fi
done

#  Note that running this script upon itself fools it
#+ into thinking it is a uuencoded file,
#+ because it contains both "begin" and "end".

#  Exercise:
#  --------
#  Modify this script to check each file for a newsgroup header,
#+ and skip to next if not found.

exit 0
