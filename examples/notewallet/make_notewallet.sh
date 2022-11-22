#!/usr/bin/env bash

cmd="../../twopasswords"

iterations="21000000"

first=$(tr -cd '[:alnum:]' < /dev/urandom | fold -w11 | head -n1)
second=$(tr -cd '[:alnum:]' < /dev/urandom | fold -w11 | head -n1)

first=$(echo $first | awk '{print toupper($0)}')
second=$(echo $second | awk '{print toupper($0)}')

echo "ITERATIONS: $iterations"
echo "FIRST:      $first"
echo "SECOND:     $second"

$cmd balance -i $iterations -p $first -s $second
