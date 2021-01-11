#!/usr/bin/bash

# INFO: this should be stored in library of some kind (no shabang!!! at the
# beginning of the library-like file

# internal library function for case changing
_caser() {
    local ret="(null)"
    while getopts 'luc' OPT;do
        case $OPT in                                      # MUTUALLY EXCLUSIVE!
            l)
                shift
                # easier with ${var^^}
                [[ $# -gt 1 ]] && ret="$(sed -nr 's/.+/\L&/p' <<< "$@")"
                ;;
            u)
                shift
                # easier with ${var,,}
                [[ $# -gt 1 ]] && ret="$(sed -nr 's/.+/\U&/p' <<< "$@")"
                ;;
            c)
                shift
                # `\b' is the word boundary(same as `\<')
                #[[ $# -gt 1 ]] && ret="$(sed -nr 's/\b./\U&/gp' <<< "$@")"
                # `\w' is a word char
                [[ $# -gt 1 ]] && ret="$(sed -nr 's/\w+/\L\u&/gp' <<< "$@")"
                ;;
        esac
    done

    shift $((OPTIND - 1))
    echo "$ret"
}

tolower () {
    echo "$(_caser -l "$@")"
}

toupper () {
    echo "$(_caser -u "$@")"
}

capitalize () {
    echo "$(_caser -c "$@")"
}

# some testing
first="putin has a SMALL dick"
second="Which Xi wants to lick."
toupper "$first" "$second"
tolower "$first" "$second"
capitalize "$first" "$second"
