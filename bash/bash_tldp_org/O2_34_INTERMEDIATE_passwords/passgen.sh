#!/usr/bin/bash

# (pseudo)random password generator
URANDOM=                                       # one bite of (pseudo)randomness
LEN=8                                                    # length of a password
NUM=2                                  # min number of digits in the a password
MAXELEMENTS=62   # max number of elements in the translation table(from int to symbol)
SUCCESS=0
FAILURE=1
PASSWD=""                                                 # password contiainer
TT=( {0..9} {a..z} {A..Z} )                                 # translation table

Read_Urandom () { # stores the last /dev/urandom byte
    URANDOM="$(od -vAn -N1 -tu4 < /dev/urandom)"
    [[ $URANDOM -gt $(( MAXELEMENTS - 1 )) ]] && Read_Urandom # infinite recursion danger
}

Read_SSLrandom () {
    URANDOM=$(( 0x$(openssl rand -hex 1) ))
    [[ $URANDOM -gt $(( MAXELEMENTS - 1 )) ]] && Read_SSLrandom # infinite recursion danger
}

testing() {
    local i
    for ((i=0; i<1000; i++)) { $1; }
}

#time testing Read_Urandom                       ---------> a little bit faster
#time testing Read_SSLrandom

# makes proper password
make_passwd() {
    local i

    PASSWD=""
    for ((i=0; i<LEN; i++)) {
        Read_Urandom
        PASSWD+=${TT[URANDOM]}
    }
}

# checks if there are at least 2 digits in a password
check_ints() {
    local ret=$SUCCESS
    local passwd="${PASSWD//[[:digit:]]/}"
    [[ ${#passwd} -gt $(( LEN - NUM ))  ]] && ret=$FAILURE
    echo $ret
}

while : ;do
    make_passwd
#    echo "internal: $PASSWD"
    [[ $(check_ints) -eq $SUCCESS ]] && break
done

echo $PASSWD
