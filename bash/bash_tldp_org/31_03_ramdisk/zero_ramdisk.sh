#!/bin/bash
# ramdisk.sh

#  A "ramdisk" is a segment of system RAM memory
#+ which acts as if it were a filesystem.
#  Its advantage is very fast access (read/write time).
#  Disadvantages: volatility, loss of data on reboot or powerdown,
#+                less RAM available to system.
#
#  Of what use is a ramdisk?
#  Keeping a large dataset, such as a table or dictionary on ramdisk,
#+ speeds up data lookup, since memory access is much faster than disk access.


E_NON_ROOT_USER=70             # Must run as root.
ROOTUSER_NAME=root
SUCCESS=0
MOUNTPT="/mnt/ramdisk"           # Create with mkdir /mnt/ramdisk.
SIZE=2000                      # 2K blocks (change as appropriate)
BLOCKSIZE=1024                 # 1K (1024 byte) block size
DEVICE="/dev/ram"               # First ram device

username="$(id -nu)"
if [[ "$username" != "$ROOTUSER_NAME" ]];then
    echo "Must be root to run <$(basename $0)>."
    exit $E_NON_ROOT_USER
fi

FOUND=0
if [[ ! -d "$MOUNTPT" ]];then
    mkdir $MOUNTPT                   #  Test whether mount point already there,
else
    FOUND=1
fi

##############################################################################
E_DD=71
dd if="/dev/zero" of=$DEVICE count=$SIZE bs=$BLOCKSIZE  # Zero out RAM device.
                                                       # Why is this necessary?
if [[ $? -ne $SUCCESS ]];then
    echo "`dd` error ocurred." &>2
    [[ $FOUND -eq 1 ]] && rm -rf "$MOUNTPT"
    exit $E_DD
fi

E_MKE2FS=72
mke2fs $DEVICE                 # Create an ext2 filesystem on it.
if [[ $? -ne $SUCCESS ]];then
    echo "`mke2df` error ocurred." &>2
    [[ $FOUND -eq 1 ]] && rm -rf "$MOUNTPT"
    exit $E_MKE2FS
fi

E_MNT=73
mount $DEVICE $MOUNTPT         # Mount it.
if [[ $? -ne $SUCCESS ]];then
    echo "`mount` error ocurred." &>2
    [[ $FOUND -eq 1 ]] && rm -rf "$MOUNTPT"
    exit $E_MNT
fi

chmod 777 $MOUNTPT             # Enables ordinary user to access ramdisk.
                               # However, must be root to unmount it.
##############################################################################
# Need to test whether above commands succeed. Could cause problems otherwise.
# Exercise: modify this script to make it safer.

echo "\"$MOUNTPT\" now available for use."
# The ramdisk is now accessible for storing files, even by an ordinary user.

#  Caution, the ramdisk is volatile, and its contents will disappear
#+ on reboot or power loss.
#  Copy anything you want saved to a regular directory.

# After reboot, run this script to again set up ramdisk.
# Remounting /mnt/ramdisk without the other steps will not work.

#  Suitably modified, this script can by invoked in /etc/rc.d/rc.local,
#+ to set up ramdisk automatically at bootup.
#  That may be appropriate on, for example, a database server.

exit 0
