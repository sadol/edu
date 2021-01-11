#!/bin/bash
# rp.sdcard.sh
# Preparing an SD card with a bootable image for the Raspberry Pi.

# $1 = imagefile name
# $2 = sdcard (device file)
# Otherwise defaults to the defaults, see below.

DEFAULTbs=4M                                 # Block size, 4 mb default.
DEFAULTif="2013-07-26-wheezy-raspbian.img"   # Commonly used distro.
DEFAULTsdcard="/dev/mmcblk0"                 # May be different. Check!
ROOTUSER_NAME=root                           # Must run as root!
E_NOTROOT=81
E_NOIMAGE=82

username=$(id -nu)                           # Who is running this script?
if [[ "$username" != "$ROOTUSER_NAME" ]];then
    echo "This script must run as root or with root privileges."
    exit $E_NOTROOT
fi

if [[ -n "$1" ]];then
    imagefile="$1"
    [[ -e "$PWD/$1" ]] && imagefile="$PWD/$1"
else
    imagefile="$PWD/$DEFAULTif"
fi

if [[ ! -e "$imagefile" ]];then
    echo "No image <$imagefile> detected." 1>&2
    exit $E_NOIMAGE
fi

if [[ -n "$2" ]];then
  sdcard="$2"
else
  sdcard="$DEFAULTsdcard"
fi

if [[ ! -e $imagefile ]];then
  echo "Image file \"$imagefile\" not found!"
  exit $E_NOIMAGE
fi

echo "Last chance to change your mind!"; echo
read -s -n1 -p "Hit a key to write $imagefile to $sdcard [Ctl-c to exit]."
echo; echo

echo "Writing $imagefile to $sdcard ..."
dd bs=$DEFAULTbs if=$imagefile of=$sdcard

exit $?

# Exercises:
# ---------
# 1) Provide additional error checking.
# 2) Have script autodetect device file for SD card (difficult!). i do not have
#    sd slot in my machine thus it is impossible to me to test it
# 3) Have script sutodetect image file (*img) in $PWD.
