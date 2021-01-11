#!/usr/bin/bash

# Masterpiece of obfuscation (for guys who like to suck their own cocks).
for ((i=1000; i<=10000; i++)) {
    for (( temp=i; ${#temp}>=2; temp=s )) {
        for ((j=0, s=0; j<${#temp}; j++)) { (( s += ${temp:j:1} )); }
    }
    (( temp == 7 )) && echo $i
}
