#!/usr/bin/bash

# use OPENSSL instead of md5sha
E_NOSSL=55
SUCCESS=0
USAGE="Usage: \`$(basename $0)' sends 6 digit pseudorandom string to \`stdout'.
\`OpenSSL' package must be installed on the system!"

which openssl > /dev/null 2>&1

if [[ $? -ne $SUCCESS ]];then
    echo "$USAGE" >&2
    exit $E_NOSSL
fi

openssl rand -hex 3                                    # 3 bytes means 6 digits
