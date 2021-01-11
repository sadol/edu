#!/bin/bash
# blot-out.sh: Erase "all" traces of a file.

#  This script overwrites a target file alternately
#+ with random bytes, then zeros before finally deleting it.
#  After that, even examining the raw disk sectors by conventional methods
#+ will not reveal the original file data.
# INFO: wiping out one file can be inefficient in case of fragmented disk space
PASSES=7         #  Number of file-shredding passes.
                 #  Increasing this slows script execution,
                 #+ especially on large target files.
BLOCKSIZE=1      #  I/O with /dev/urandom requires unit block size,
                 #+ otherwise you get weird results.
E_BADARGS=70     #  Various error exit codes.
E_NOT_FOUND=71
E_CHANGED_MIND=72

if [[ -z "$1" ]];then   # No filename specified.
    echo "Usage: $(basename $0) filename" 1>&2
    exit $E_BADARGS
fi

file=$1

if [[ ! -e "$file" ]];then
    echo "File \"$file\" not found." 1>&2
    exit $E_NOT_FOUND
fi

echo; echo -n "Are you absolutely sure you want to blot out \"$file\" (y/n)? "
read answer
case "$answer" in
    [nN])
        echo "Changed your mind, huh?"
        exit $E_CHANGED_MIND
        ;;
    *)
        echo "Blotting out file \"$file\".";;
esac


flength=$(ls -l "$file" | awk '{print $5}')  # Field 5 is file length.
pass_count=1

chmod u+w "$file"   # Allow overwriting/deleting the file.

echo

while [[ "$pass_count" -le "$PASSES" ]];do
    echo "Pass #$pass_count"
    sync         # Flush buffers.
    dd if=/dev/urandom of=$file bs=$BLOCKSIZE count=$flength iflag=fullblock
                # Fill with random bytes.
    sync         # Flush buffers again.
    dd if=/dev/zero of=$file bs=$BLOCKSIZE count=$flength iflag=fullblock
                # Fill with zeros.
    sync         # Flush buffers yet again.
    ((pass_count++))
    echo
done

sync         # Flush buffers.
# do not leave 0-ied space; white noise is better
dd if=/dev/urandom of=$file bs=$BLOCKSIZE count=$flength iflag=fullblock
sync

rm -f $file    # Finally, delete scrambled and shredded file.
sync           # Flush buffers a final time.

# NOOOO shit ?
echo "File \"$file\" blotted out and deleted."; echo

# OOOOOOOOOOOOOOOOOOOOOOOOOOOR just use `wipe' for example:
wipe $file

exit 0

#  This is a fairly secure, if inefficient and slow method
#+ of thoroughly "shredding" a file.
#  The "shred" command, part of the GNU "fileutils" package,
#+ does the same thing, although more efficiently.

#  The file cannot not be "undeleted" or retrieved by normal methods.
#  However . . .
#+ this simple method would *not* likely withstand
#+ sophisticated forensic analysis.

#  This script may not play well with a journaled file system.
#  Exercise (difficult): Fix it so it does.



#  Tom Vier's "wipe" file-deletion package does a much more thorough job
#+ of file shredding than this simple script.
#     http://www.ibiblio.org/pub/Linux/utils/file/wipe-2.0.0.tar.bz2

#  For an in-depth analysis on the topic of file deletion and security,
#+ see Peter Gutmann's paper,
#+     "Secure Deletion of Data From Magnetic and Solid-State Memory".
#       http://www.cs.auckland.ac.nz/~pgut001/pubs/secure_del.html
