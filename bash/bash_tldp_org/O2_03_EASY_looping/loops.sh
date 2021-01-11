#!/bin/bash
PLANETS=( Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune Pluto )

sentinel=0
while [[ $sentinel -lt ${#PLANETS[@]} ]];do
  echo ${PLANETS[$sentinel]}  # Each planet on a separate line.
  (( sentinel++ ))
done

echo

sentinel=0
until [[ $sentinel -gt ${#PLANETS[@]} ]];do
  echo ${PLANETS[$sentinel]}  # Each planet on a separate line.
  (( sentinel++ ))
done

echo; echo "Whoops! Pluto is no longer a planet!"

exit 0
