#! /bin/bash
# is-spammer.sh: Identifying spam domains

# $Id: is-spammer, v 1.4 2004/09/01 19:37:52 mszick Exp $
# Above line is RCS ID info.
#
#  This is a simplified version of the "is_spammer.bash
#+ script in the Contributed Scripts appendix.

# is-spammer <domain.name>

# Uses an external program: 'dig'
# Tested with version: 9.2.4rc5

# Uses functions.
# Uses IFS to parse strings by assignment into arrays.
# And even does something useful: checks e-mail blacklists.

# Use the domain.name(s) from the text body:
# http://www.good_stuff.spammer.biz/just_ignore_everything_else
#                       ^^^^^^^^^^^
# Or the domain.name(s) from any e-mail address:
# Really_Good_Offer@spammer.biz
#
# as the only argument to this script.
#(PS: have your Inet connection running)
#
# So, to invoke this script in the above two instances:
#       is-spammer.sh spammer.biz


# Whitespace == :Space:Tab:Line Feed:Carriage Return:
WSP_IFS=$'\x20'$'\x09'$'\x0A'$'\x0D'

# No Whitespace == Line Feed:Carriage Return
No_WSP=$'\x0A'$'\x0D'

# Field separator for dotted decimal ip addresses
ADR_IFS=${No_WSP}'.'

# Get the dns text resource record.
# get_txt <error_code> <list_query>
get_txt() {
    # Parse $1 by assignment at the dots.
    local -a dns
    IFS=$ADR_IFS
    dns=( $1 )
    IFS=$WSP_IFS
    if [[ "${dns[0]}" == '127' ]];then
        # See if there is a reason.
        echo $(dig +short $2 -t txt)
    fi
}

# Get the dns address resource record.
# chk_adr <rev_dns> <list_server>
chk_adr() {
    local reply
    local server
    local reason

    server=${1}${2}
    reply=$(dig +short ${server})

    # If reply might be an error code . . .
    if [[ ${#reply} -gt 6 ]];then
        reason=$(get_txt ${reply} ${server})
        reason=${reason:-${reply}}
    fi
    echo ${reason:-' not blacklisted.'}
}

# slightly upgraded stackoverflow solution
chk_net () {
    local path="/sys/class/net/"

    for interface in $(ls "$path" | grep -v lo);do
        if [[ $(cat "$path""$interface"/carrier 2>/dev/null) = "1" ]];then
            echo "OnLine"
            return
        fi
    done

    echo "OffLine"
}

## -----------------main()------------------------------
## -----------------------------------------------------
E_OFFLINE=78
OFFLINEMSG="User <$USER> must be online to run this script."
LINE_STATE="$(chk_net)"
if [[ "$LINE_STATE" = "OffLine" ]];then
    echo "$OFFLINEMSG" 1>&2
    exit $E_OFFLINE
fi

E_BADARG=77
DEFAULT_BLHS="sbl-xbl.spamhaus.org relays.ordb.org bl.spamcop.net \
cbl.abuseat.org list.dsbl.org unconfirmed.dsbl.org multihop.dsbl.org"
USAGEMSG="Usage: $(basename $0) -s <spam-site> [-b BLH-space-delimited-list]"
SPAMMER=""
BLHS=""

while getopts ":s:b:" OPT;do
    case "$OPT" in
        s)
            SPAMMER="$OPTARG"
            shift
            ;;
        b)
            max=$(($# - 1))
            for (( i=0 ;i<$max ;i++ )){
                BLHS="$BLHS $OPTARG"
                shift
            }
            ;;
        :)
            echo "Unknown option <$OPTARG>." 1>&2
            echo "$USAGEMSG" 1>&2
            exit $E_BADARG
            ;;
        \?)
            echo "$USAGEMSG" 1>&2
            exit $E_BADARG
            ;;
    esac
done

[[ "$BLHS" = "" ]] && BLHS="$DEFAULT_BLHS"

ERRORMSG="Argument $SPAMMER should have form of .+\..{1,3}"
sedcheck=$(sed -rn '/^.+\..{1,3}$/ p' <<< "$SPAMMER")
if [[ $# -ne 1 || ${#sedcheck} -eq 0 ]];then
    echo "$ERRORMSG" 1>&2
    exit $E_BADARG
fi

# Need to get the IP address from the name.
echo 'Get address of: '$1
ip_adr=$(dig +short $1)
dns_reply=${ip_adr:-' no answer '}
echo ' Found address: '${dns_reply}

# A valid reply is at least 4 digits plus 3 dots.
if [[ ${#ip_adr} -gt 6 ]];then
    echo
    declare query

    # Parse by assignment at the dots.
    declare -a dns
    IFS=$ADR_IFS
    dns=( ${ip_adr} )
    IFS=$WSP_IFS

    # Reorder octets into dns query order.
    rev_dns="${dns[3]}"'.'"${dns[2]}"'.'"${dns[1]}"'.'"${dns[0]}"'.'

    for blh in $BLHS;do
        echo -n "$blh says:"
        echo $(chk_adr ${rev_dns} $blh)
    done
else
    echo
    echo 'Could not use that address.'
fi

exit 0

# Exercises:
# --------

# 1) Check arguments to script,
#    and exit with appropriate error message if necessary.

# 2) Check if on-line at invocation of script,
#    and exit with appropriate error message if necessary.

# 3) Substitute generic variables for "hard-coded" BHL domains.

# 4) Set a time-out for the script using the "+time=" option
#    to the 'dig' command.
