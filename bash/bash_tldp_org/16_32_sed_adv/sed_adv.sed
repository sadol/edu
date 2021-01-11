# check if there is onliner comment in a line (only comment in the line)

# non extended version of sed -> impressive picket fence if i may say so myself
# s/\(.*\)\(\/\*\)\(.*\)\(\*\/\)\(.*\)/\1 \5/
# extended version of sed
s/(.*)(\/\*)(.*)(\*\/)(.*)/\1 \5/
# remove single spaced lines as a possible result of the previous action
/^ $/ d
# remove leading spaces
s/^ +(.+)/\1/
# remove trailing spaces
s/(.+) +$/\1/
# then check if the are multiliners
/\/\*/, /\*\// d
p
