#!/bin/bash
# monthlypmt.sh: Calculates monthly payment on a mortgage.

#  This is a modification of code in the
#+ "mcalc" (mortgage calculator) package,
#+ by Jeff Schmidt
#+ and
#+ Mendel Cooper (yours truly, the ABS Guide author).
#   http://www.ibiblio.org/pub/Linux/apps/financial/mcalc-1.6.tar.gz

echo
echo "Given the principal, interest rate, and term of a mortgage,"
echo "calculate the monthly payment."

bottom=1.0

echo
echo -n "Enter principal : "
read principal
[[ ! ${principal#*,} = $principal ]] && principal=$(sed -rn 's/,//g p' <<< "$principal")
echo -n "Enter interest rate :  "  # If 12%, enter "12", not ".12".
read interest_r
# [[ ! $(sed -nr '/^.*\..+$/p' <<< "$interest_r") ]] && interest_r=$(bc <<< "scale=9; $interest_r/100.0")
[[ ! $(sed -nr '/^.*\..+$/p' <<< "$interest_r") ]] && interest_r=$(dc <<< "9 k $interest_r 100.0 / p")
echo -n "Enter term (months) "
read term
interest_rate=$(dc <<< "9 k $interest_r 12 / 1 + p")

#           Standard formula for figuring interest.
top=$(bc <<< "scale=9; $principal*$interest_rate^$term")
#echo; echo "Please be patient. This may take a while."

months=$((term - 1))
# ====================================================================
#for ((x=$months; x > 0; x--));do
#    bot=$(bc <<< "scale=9; $interest_rate^$x")
#    bottom=$(bc <<< "scale=9; $bottom+$bot")
#    #  bottom = $(($bottom + $bot))
#done
# ====================================================================

# --------------------------------------------------------------------
#  Rick Boivie pointed out a more efficient implementation
#+ of the above loop, which decreases computation time by 2/3.

# for ((x=1; x <= $months; x++))
# do
#   bottom=$(echo "scale=9; $bottom * $interest_rate + 1" | bc)
# done


#  And then he came up with an even more efficient alternative,
#+ one that cuts down the run time by about 95%!

# bottom=`{
#     echo "scale=9; bottom=$bottom; interest_rate=$interest_rate"
#     for ((x=1; x <= $months; x++))
#     do
#          echo 'bottom = bottom * interest_rate + 1'
#     done
#     echo 'bottom'
#     } | bc`       # Embeds a 'for loop' within command substitution.
# --------------------------------------------------------------------------
#  On the other hand, Frank Wang suggests:
bottom=$(bc <<< "scale=9; ($interest_rate^$term-1)/($interest_rate-1)")
#  Because . . .
#  The algorithm behind the loop
#+ is actually a sum of geometric proportion series.
#  The sum formula is e0(1-q^n)/(1-q),
#+ where e0 is the first element and q=e(n+1)/e(n)
#+ and n is the number of elements.
# --------------------------------------------------------------------------


# let "payment = $top/$bottom"
payment=$(dc <<< "2 k $top $bottom / p")
# Use two decimal places for dollars and cents.

echo
echo "monthly payment = \$$payment"  # Echo a dollar sign in front of amount.
echo

echo "----------------------- AMORTIZATION TABLE---------------------------------"
echo "| Mth. | Payment | Interest |    Loan    | Interest Paid | Principal Paid |"
echo "---------------------------------------------------------------------------"
prev_loan=$principal
loan=
for ((x=1; x<=term ;x++ ));do
    loan=$prev_loan
    interest_paid=$(dc <<< "9 k $loan $interest_r * p")
    principal_paid=$(dc <<< "9 k $payment $interest_paid - p")
    prev_loan=$(dc <<< "9 k $loan $principal_paid - p")
    printf "|%6d|" $x
    printf "%9.2f|" $(sed -nr 's/\./,/p' <<< $payment)
    printf "%10.2f|" $(sed -nr 's/\./,/p' <<< $interest_r)
    printf "%12.2f|" $(sed -nr 's/\./,/p' <<< $loan)
    printf "%15.2f|" $(sed -nr 's/\./,/p' <<< $interest_paid)
    printf "%16.2f|\n" $(sed -nr 's/\./,/p' <<< $principal_paid)
done
echo "---------------------------------------------------------------------------"
exit 0

# Exercises:
#   1) Filter input to permit commas in principal amount.
#   2) Filter input to permit interest to be entered as percent or decimal.
#   3) If you are really ambitious,
#+     expand this script to print complete amortization tables.
