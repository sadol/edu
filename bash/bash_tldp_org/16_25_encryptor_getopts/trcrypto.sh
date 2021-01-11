#!/bin/bash
# crypto-quote.sh: Encrypt quotes

#  Will encrypt famous quotes in a simple monoalphabetic substitution.
#  The result is similar to the "Crypto Quote" puzzles
#+ seen in the Op Ed pages of the Sunday paper.

# The "KEY" is nothing more than a scrambled alphabet.
# Changing the "key" changes the encryption.
#  This simple-minded cipher can be broken by an average 12-year old
#+ using only pencil and paper.

E_WRONGARGS=85
USAGEMSG="Usage: $(basename $0) [-d] [-k key] [filename]"
FILENAME=""
KEY="ETAOINSHRDLUBCFGJMQPVWZYXK"
DECODE=0
TEST=0

while getopts ":dtk:" OPT; do
    case "$OPT" in
        d)
            DECODE=1
            ;;
        k)
            KEY="$OPTARG"
            ;;
        t)  # for mark twain testing
            TEST=1
            ;;
        :)
            echo "Unknown option <$OPTARG>." 1>&2
            exit $E_WRONGARG
            ;;
        \?)
            echo "$USAGEMSG" 1>&2
            exit $E_WRONGARG
            ;;
    esac
done

shift $((OPTIND - 1))

if [[ -n "$1" ]]; then
    FILENAME="$1"
    if [[ ! -f "$FILENAME" ]]; then
        echo "File name argument <$FILENAME> not a regular file." 1>&2
        echo "$USAGEMSG" 1>&2
        exit $E_WRONGARGS
    fi
fi

case $DECODE in
    0)
        # The 'cat "$@"' construction gets input either from stdin or from files.
        # If using stdin, terminate input with a Control-D.
        # Otherwise, specify filename as command-line parameter.

        cat "$@" | tr [:lower:] [:upper:] | tr [:upper:] "$KEY"
        #        |  to uppercase          |     encrypt
        # Will work on lowercase, uppercase, or mixed-case quotes.
        # Passes non-alphabetic characters through unchanged.
        ;;
    1)
        # To reverse the encryption:
        # THIS DOES NOT WORK -> character class error ;( , locale related??????
        #cat "$@" | tr "$KEY" [:upper:]
        # this works fine
        cat "$@" | tr "$KEY" "A-Z"
        ;;
esac

# Try this script with something like:
# "Nothing so needs reforming as other people's habits."
# --Mark Twain
#
# Output is:
# "CFPHRCS QF CIIOQ MINFMBRCS EQ FPHIM GIFGUI'Q HETRPQ."
# --BEML PZERC
exit $?

#  Exercise:
#  --------
#  Modify the script so that it will either encrypt or decrypt,
#+ depending on command-line argument(s).
