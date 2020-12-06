#!/bin/bash

count=0

while [ "$count" -lt 30 ]; do
    sleep 1000 &
    count=$(( count + 1 ))
    echo "$count"
done

wait