#!/bin/bash
# quote-fetch.sh: Download a stock quote.

# slightly upgraded stackoverflow solution
chk_net () {
    local path="/sys/class/net/"

    for interface in $(ls "$path" | grep -v lo);do
        if [[ "$(cat ${path}${interface}/carrier 2>/dev/null)" = "1" ]];then
            echo "OnLine"
            return
        fi
    done

    echo "OffLine"
}

E_OFFLINE=78
OFFLINEMSG="User <$USER> must be online to run this script."
LINE_STATE="$(chk_net)"
if [[ "$LINE_STATE" = "OffLine" ]];then
    echo "$OFFLINEMSG" 1>&2
    exit $E_OFFLINE
fi

E_NOPARAMS=86

if [[ -z "$1" ]];then  # Must specify a stock (symbol) to fetch.
    echo "Usage: $(basename $0) stock-symbol"
    exit $E_NOPARAMS
fi

stock_symbol=$1

file_suffix=".html"
# Fetches an HTML file, so name it appropriately.
URL1='https://finance.yahoo.com/quote/'
URL2='&.tsrc=fin-srch'
# Yahoo finance board, with stock query suffix.

# -----------------------------------------------------------
wget -O ${stock_symbol}${file_suffix} "${URL1}${stock_symbol}${URL2}"
# -----------------------------------------------------------

URL3='https://twojapogoda.pl/prognoza-polska/'
voivodeship="$2"
city="$3"

#
wget -O $city$file_suffix "$URL3$voivodeship-$city"
# To look up stuff on http://search.yahoo.com:
# -----------------------------------------------------------
# URL="http://search.yahoo.com/search?fr=ush-news&p=${query}"
# wget -O "$savefilename" "${URL}"
# -----------------------------------------------------------
# Saves a list of relevant URLs.

exit $?

# Exercises:
# ---------
#
# 1) Add a test to ensure the user running the script is on-line.
#    (Hint: parse the output of 'ps -ax' for "ppp" or "connect."
#
# 2) Modify this script to fetch the local weather report,
#+   taking the user's zip code as an argument.
