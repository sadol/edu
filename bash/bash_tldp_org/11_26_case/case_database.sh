#!/bin/bash

# Crude address database
EXIT="e"

show_menu () {
    clear # Clear the screen.

    echo "          Contact List"
    echo "          ------- ----"
    echo "Choose one of the following persons:"
    echo
    echo "[E]vans, Roland"
    echo "[J]ones, Mildred"
    echo "[S]mith, Julie"
    echo "[Z]ane, Morris"
    echo
}

ask_to_continue () {
    local conti=""
    echo "------------------------"
    echo "Hit <ENTER> to continue or <"$EXIT"> to exit"
    read conti
    case "$conti" in
        "$EXIT" )
            echo
            echo "Bye!"
            exit   # not very robust (second exit point from the script) but simple
            ;;
        * )    # continue
            ;;
    esac
}

show_personal_entry () {
    case "$person" in
    # Note variable is quoted.

    "E" | "e" )
        echo
        echo "Roland Evans"
        echo "4321 Flash Dr."
        echo "Hardscrabble, CO 80753"
        echo "(303) 734-9874"
        echo "(303) 734-9892 fax"
        echo "revans@zzy.net"
        echo "Business partner & old friend"
        ;;

    "J" | "j" )
        echo
        echo "Mildred Jones"
        echo "249 E. 7th St., Apt. 19"
        echo "New York, NY 10009"
        echo "(212) 533-2814"
        echo "(212) 533-9972 fax"
        echo "milliej@loisaida.com"
        echo "Ex-girlfriend"
        echo "Birthday: Feb. 11"
        ;;

    # Add info for Smith & Zane later.
    * )
        # Default option.
        # Empty input (hitting RETURN) fits here, too.
        echo
        echo "Not yet in database."
        ;;
    esac

    echo
    ask_to_continue
}

for (( ; ; )) {
    show_menu
    read person
    show_personal_entry
}

#  Exercise:
#  --------
#  Change the script so it accepts multiple inputs,
#+ instead of terminating after displaying just one address.

exit 0
