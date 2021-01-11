#!/usr/bin/bash

# nonintercative script ;returns 0 if process is still runnig xor 1 otherwise;

RUNNIG=0
NOTRUNNIG=1
SUCCESS=0
USAGE="Usage: $(basename $0) -p <PID> -i <interval-in-sec>."
E_ARGS=66
INTERVAL=
PROCID=

while optargs ':p:i:' OPT; do
    case $OPT in
        p) PROCID="$OPTARG";;
        i) INTERVAL="$OPTARG";;
        :)
            echo "Unknown option argument: <$OPTARG>." >&2
            echo "$USAGE" >&2
            exit $E_ARGS
            ;;
        \?)
            echo "$USAGE" >&2
            exit $E_ARGS
            ;;
    esac
done

shift $(( OPTIND - 1 ))

while : ;do
    ps r -"$PROCID" > /dev/null 2>&1
    [[ $? -eq $SUCCESS ]] && echo $RUNNIG || echo $NOTRUNNIG
    sleep "$INTERVAL"
done
