#!/bin/bash
# Creating a swap file.

#  A swap file provides a temporary storage cache
#+ which helps speed up certain filesystem operations.

ROOT_UID=0         # Root has $UID 0.
E_WRONG_USER=85    # Not root?

FILE=/swap
BLOCKSIZE=1024
MINBLOCKS=40
SUCCESS=0


# This script must be run as root.
if [[ "$UID" -ne "$ROOT_UID" ]];then
    echo; echo "You must be root to run this script."; echo
    exit $E_WRONG_USER
fi

blocks=${1:-$MINBLOCKS}          #  Set to default of 40 blocks,
                                 #+ if nothing specified on command-line.

[[ "$blocks" -lt $MINBLOCKS ]] && blocks=$MINBLOCKS

######################################################################
echo "Creating swap file of size $blocks blocks (KB)."
E_DD=86
dd if=/dev/zero of=$FILE bs=$BLOCKSIZE count=$blocks  # Zero out file.
if [[ $? -ne $SUCCESS ]];then
    echo "`dd` error ocurred." &>2
    exit $E_DD
fi

E_SWAP=87
mkswap $FILE $blocks             # Designate it a swap file.
if [[ $? -ne $SUCCESS ]];then
    echo "`swap` error ocurred." &>2
    rm -rf "$FILE"
    exit $E_SWAP
if

E_SWAPON=88
swapon $FILE                     # Activate swap file.
retcode=$?                       # Everything worked?
if [[ $retcode -ne $SUCCESS ]];then
    echo "`swapon` error ocurred." &>2
    rm -rf "$FILE"
    exit $E_SWAPON
fi
######################################################################

#  Exercise:
#  Rewrite the above block of code so that if it does not execute
#+ successfully, then:
#    1) an error message is echoed to stderr,
#    2) all temporary files are cleaned up, and
#    3) the script exits in an orderly fashion with an
#+      appropriate error code.

echo "Swap file created and activated."

exit $retcode
