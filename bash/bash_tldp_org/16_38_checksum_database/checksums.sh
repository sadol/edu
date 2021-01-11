#!/bin/bash
# file-integrity.sh: Checking whether files in a given directory
#                    have been tampered with.

E_DIR_NOMATCH=80
E_BAD_DBFILE=81
E_BAD_CHECKFILE=82
E_MISSING_FILES=83

# Filename for storing records (database file).
dbfile="File_record.sha512"
dbcheckfile="DB_record.sha512"
directory=""
sedscript='/'"$dbfile"'/ d
           /^ *$/ d
           p
'

set_up_database () {
    # Write directory name to first line of file.
    echo ""$directory"" > "$dbfile"
    # Append sha512sum checksums and filenames.
    sha512sum "$directory"/* >> "$dbfile" 2> /dev/null
    # except checksums database itself
    sed -i -n "$sedscript" "$dbfile"
    # Append sha512sum checksum to the dedicated file
    sha512sum "$dbfile" > "$dbcheckfile" 2> /dev/null
}

check_database () {
    local n=0
    local filename
    local checksum

    # ------------------------------------------- #
    #  This file check should be unnecessary,
    #+ but better safe than sorry.

    if [[ ! -r "$dbfile" ]]; then
        echo "Unable to read checksum database file!" 1>&2
        exit $E_BAD_DBFILE
    fi

    if [[ ! -r "$dbcheckfile" ]]; then
        echo "Unable to read control checksum file!" 1>&2
        exit $E_BAD_CHECKFILE
    fi

    # ------------------------------------------- #

    while read record; do
        if [[ $n -eq 0 ]];then
            directory_checked="${record}"
            if [[ "$directory_checked" != "$directory" ]];then
                # Tried to use file for a different directory.
                echo "Directories do not match up!" 1>&2
                exit $E_DIR_NOMATCH
            fi
        else                                              # Not directory name.
            #  sha512sum writes records backwards,
            #+ checksum first, then filename.
            filename="$( awk '{ print $2 }' <<< "$record" )"
            checksum=$( sha512sum "${filename}" )

            if [[ "${record}" = "${checksum}" ]];then
                echo "${filename} unchanged."
            else
                echo "${filename} : CHECKSUM ERROR!"
            fi
        fi

        (( n++ ))
    done <"$dbfile"       # Read from checksum database file.
}

check_checkbase() {
    local checksum=""
    local filename=""
    local record=""
    # only one record to read
    read record < "$dbcheckfile"
    filename="$( awk '{ print $2 }' <<< "$record" )"
    checksum=$( sha512sum "${filename}" )
    if [[ "${record}" = "${checksum}" ]];then
        echo "${filename} unchanged."
    else
        echo "${filename} : CHECKSUM ERROR!"
    fi
}

# =================================================== #
# main ()

if [[ -z  "$1" ]];then
    directory="$PWD"      #  If not specified,
else                    #+ use current working directory.
    directory="$1"
fi

clear                   # Clear screen.
echo " Running file integrity check on $directory"
echo

# ------------------------------------------------------------------ #
if [[ ! -r "$dbfile" && -r "$dbcheckfile" || -r "$dbfile" && ! -r "$dbcheckfile" ]];then
    echo "Checksums files tampering detected!"
    echo "Remove "$dbfile" xor "$dbcheckfile" and run the script again."
    exit $E_MISSING_FILES
elif [[ ! ( -r "$dbfile" || -r "$dbcheckfile" ) ]];then
    echo "Setting up database file, \""$directory"/"$dbfile"\"."; echo
    echo "Setting up check file, \""$directory"/"$dbcheckfile"\"."; echo
    set_up_database
else # both databases are present
    check_database          # Do the actual work.
    echo
    check_checkbase          # Do the actual work.
fi

#  You may wish to redirect the stdout of this script to a file,
#+ especially if the directory checked has many files in it.

exit 0

#  For a much more thorough file integrity check,
#+ consider the "Tripwire" package,
#+ http://sourceforge.net/projects/tripwire/.
