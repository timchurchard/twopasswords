#!/bin/bash

if [ -z $PASSWORD ]; then
	export PASSWORD=$(goxkcdpwgen -d - -c -n 10)
fi

if [ -z $BITS ]; then
	export BITS=128
fi

if [ -z $SCRIPT ]; then
	export SCRIPT=p2wpkh
fi

export ITS=$(shuf -i 101101-212121 -n 1)

twopasswords address -b $BITS -i $ITS -n 0 -p $PASSWORD -script $SCRIPT -s ${ITS} | grep Mnemonic
twopasswords address -b $BITS -i $ITS -n 0 -p $PASSWORD -script $SCRIPT -s ${ITS} | grep address
