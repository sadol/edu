#!/bin/bash

MAX=10000

#nr=

for ((nr=1; nr<MAX; nr++)) {
    t1=$(( nr % 5 ))
    [[ "$t1" -ne 3 ]] && continue

    t2=$(( nr % 7 ))
    [[ "$t2" -ne 4 ]] && continue                               # modulo 5 == 3

    t3=$(( nr % 9 ))
    [[ "$t3" -ne 5 ]] && continue              # modulo 5 == 3 && modulo 7 == 4

    break                     # modulo 5 == 3 && modulo 7 == 4 && modulo 9 == 5
}

echo "Number = $nr"
exit 0
