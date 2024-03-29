#!/bin/bash
#  cvt.sh:
#  Converts all the MacPaint image files in a directory to "pbm" format.

#  Uses the "macptopbm" binary from the "netpbm" package,
#+ which is maintained by Brian Henderson (bryanh@giraffe-data.com).
#  Netpbm is a standard part of most Linux distros.

#OPERATION=macptopbm
OPERATION=cp
SUFFIX=pbm          # New filename suffix.
OLDSUFFIX=mac

if [[ -n "$1" ]];then
    directory=$1      # If directory name given as a script argument...
else
    directory=$PWD    # Otherwise use current working directory.
fi

#  Assumes all files in the target directory are MacPaint image files,
#+ with a ".mac" filename suffix.
# INFO: beautiful oneliner
#       --->                                 ${file/%${OLDSUFFIX}/${NEWSUFFIX}}

for file in $directory/* ;do    # Filename globbing.
    if [[ $file =~ ^.*${OLDSUFFIX}$ ]]; then
        filename=${file%.$OLDSUFFIX}      #  Strip ".mac" suffix off filename
                                #+ ('.*c' matches everything
                    #+ between '.' and 'c', inclusive).
        $OPERATION $file "$filename.$SUFFIX"
                                # Redirect conversion to new filename.
        rm -f $file               # Delete original files after converting.
        echo "$filename.$SUFFIX"  # Log what is happening to stdout.
    fi
done

exit 0

# Exercise:
# --------
#  As it stands, this script converts *all* the files in the current
#+ working directory.
#  Modify it to work *only* on files with a ".mac" suffix.
