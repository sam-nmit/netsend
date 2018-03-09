#!/bin/bash

REMOTE_HOST="192.168.0.108"
PORT=5000

if [ ! -f "netsend" ]
then
    echo "netsend not found."
    exit 1
fi

if [ -z $2 ]
then
    REMOTE_HOST=$2
fi

PIPE=$(<&0)

case $1 in 
"web")
    ./netsend -w $PORT $PIPE
    ;;
"wget")
    curl -s "$REMOTE_HOST:$PORT" >&1
    ;;
"put")
    ./netsend -s "$REMOTE_HOST:$PORT" $PIPE
    ;;
"get")
    ./netsend -l $PORT >&1
    ;;
esac
