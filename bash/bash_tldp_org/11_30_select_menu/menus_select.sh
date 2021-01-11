#!/bin/bash
OLD_PS3="$PS3"
PS3='Choose your favorite vegetable: ' # Sets the prompt string.
                                       # Otherwise it defaults to #? .
veggies=("beans" "carrots" "potatoes" "onions" "rutabagas")

echo

select vegetable in ${veggies[@]};do
    echo
    if [[ -z $vegetable ]];then
        echo "You are liar! There is no such veggie!"
    else
        echo "Your favorite veggie is $vegetable."
    fi
    echo
    break  # What happens if there is no 'break' here?
done

PS3="$OLD_PS3"
exit

# Exercise:
# --------
#  Fix this script to accept user input not specified in
#+ the "select" statement.
#  For example, if the user inputs "peas,"
#+ the script would respond "Sorry. That is not on the menu."
