#!/usr/bin/bash

USAGE="USAGE(as root): $(basename $0) <user_name>."
# mounting partitions only as the root user
E_ROOT=55
if [[ $(id -u) -ne 0 ]];then
    echo "You must be ROOT to use this script."
    exit $E_ROOT
fi

E_USER=56
if [[ -z $1 ]];then
    echo "$USAGE"
    exit $E_USER
else
    USERNAME="$1"
    TARNAME="${USERNAME}_tree"
    TARBALL="${TARNAME}.tar.gz"
fi

HOMEDIR="$(sed -nr 's/^'"$USERNAME"':.*:(.+):.+$/\1/p' /etc/passwd)"
if [[ -z "$HOMEDIR" ]];then
    echo "User <$USERNAME> unknown."
    echo "$USAGE"
    exit $E_USER
fi

ls -R ${HOMEDIR} > ${TARNAME}
tar czf ${TARBALL} ${TARNAME}
rm -rf ${TARNAME}

# =================check if usb drive is present in the system=================

# info file of connected disk drives (may vary from distro to distro I suppose)
DISKSTATS="/proc/diskstats"
SEDSCRIPT='s/^   //; s/ +/ /g; s/^[^[:space:]]+ [^[:space:]]+ ([^[:space:]]+) .*$/\1/p'
# initial list of disks phisically connected to the system a.k.a. SED TIME!!!
DRIVES_LIST_INIT=( $(sed -nr -e "$SEDSCRIPT" "$DISKSTATS") )
len_init=${#DRIVES_LIST_INIT[@]}

echo -n "Put the usb drive into one of the usb slots and hit ENTER:"
read
sleep 4                          # give it a chance to update the list of disks

# list of disks phisically connected to the system (not neccesarily mounted)
DRIVES_LIST_USB=( $(sed -nr -e "$SEDSCRIPT" "$DISKSTATS") )
len_usb=${#DRIVES_LIST_USB[@]}

declare -a NEW_DRIVES_LIST
new_id=0  # number of elements of the list above
found=0   # found flag
for ((i=0; i<len_usb; i++)) {
    found=0
    for ((j=0; j<len_init; j++)) {
        if [[ ${DRIVES_LIST_USB[i]} = ${DRIVES_LIST_INIT[j]} ]];then
            found=1
            break
        fi
    }

    if [[ $found -eq 0 ]];then
        NEW_DRIVES_LIST[$new_id]=${DRIVES_LIST_USB[i]}
        (( new_id++ ))
    fi
}

E_NO_USB=57
if [[ ${#NEW_DRIVES_LIST[@]} -eq 0 ]];then
    echo "No new usb drive detected!!!"
    exit $E_NO_USB
fi

# ==============mount newly detected partition and check it====================

# create temporary mounting point for the new usb partition
SUBDIR="/opt/usb${RANDOM}"
mkdir "$SUBDIR"
E_MOUNT=58
MOUNTED=0
# try to mount first usb partition in the `/opt/$RANDOM' directory
for diskid in ${NEW_DRIVES_LIST[*]}; do
    if [[ ${diskid:${#diskid}-1} =~ [[:digit:]] ]];then
        mount /dev/$diskid "$SUBDIR"
        if [[ ! $? ]];then                                             # NOT ok
            echo "There is some problem with mounting </dev/$diskid> into <$SUBDIR>."
            exit $E_MOUNT
        fi
        MOUNTED=1
        break
    fi
done

if [[ $MOUNTED -eq 0 ]];then     # there is no partition ready on the usb drive
    echo "No ready partition on the usb drive."
    exit $E_MOUNT
fi

# =====================copy file===============================================
E_COPY=59
cp "$TARBALL" "$SUBDIR"
if [[ ! $? ]];then                                             # NOT ok
    echo "There is a problem with copying file <$TARBALL> to <$SUBDIR>."
    exit $E_COPY
fi

# =====================umount usb drive========================================
umount "$SUBDIR"
if [[ ! $? ]];then                                             # NOT ok
    echo "There is a problem with umount of <$SUBDIR>."
    exit $E_MOUNT
fi
rm -rf "$SUBDIR"
if [[ ! $? ]];then                                             # NOT ok
    echo "There is a problem with removing <$SUBDIR>."
    exit $E_MOUNT
fi

echo; echo "DONE."; echo
# =============================================================================
