#!/bin/bash
# Sanity tests to protect from breaking changes to the scripts

NEW_CMD="./twopasswords"
OLD_CMD="./run_python_twop.sh"

for lcmd in $OLD_CMD $NEW_CMD; do
  out=$($lcmd seed --password apple8)
  if [[ $out == *"b930218bd7b1b6620ebc7bf0c225a76e5c3780411ea104cc83bd40d57f6d6f6d"* ]]; then
    echo "$lcmd - sanity - 1 - seed password (apple8) hex = OK"
  fi
  if [[ $out == *"rice library glass quarter bread country depend month valid bachelor hawk syrup sell scan afford stage age goat urban like fitness repeat sad poverty"* ]]; then
    echo "$lcmd - sanity - 2 - seed password (apple8) seed = OK"
  fi

  out=$($lcmd seed --password apple8 --iterations 12345)
  if [[ $out == *"df43c0019f4c8a04132779cc871b2645b0107f8452770ed4505976e0269efa03"* ]]; then
    echo "$lcmd - sanity - 3 - seed password (apple8,12345) hex = OK"
  fi
  if [[ $out == *"tent bulk about direct silly acoustic erode upset smart decide sister merge absurd divert bacon exclude attract penalty bind universe act exhaust trend improve"* ]]; then
    echo "$lcmd - sanity - 4 - seed password (apple8,12345) seed = OK"
  fi

  out=$($lcmd address --password apple8 --second orange2 --mode python)
  if [[ $out == *"bc1q6nhhvxkygue0pg89ac96ax2dcphkazmgydc8rc"* ]]; then
    echo "$lcmd - sanity - 5 python - address (apple8,orange2) address = OK"
  fi

  out=$($lcmd address --password apple8 --second orange2 --num 0 --mode python)
  if [[ $out == *"bc1q6nhhvxkygue0pg89ac96ax2dcphkazmgydc8rc"* ]]; then
    echo "$lcmd - sanity - 6 python - address (apple8,orange2,0) address = OK"
  fi

  out=$($lcmd address --password apple8 --second orange2 --num 1 --mode python)
  if [[ $out == *"bc1q5xm0fv5v8dmv6kdz42mz6auakp5xstqsc9r5vq"* ]]; then
    echo "$lcmd - sanity - 7 python - address (apple8,orange2,1) address = OK"
  fi

  out=$($lcmd address --password apple8 --second orange2 --iterations 12345 --mode python)
  if [[ $out == *"bc1qzsjra9qf52fu8mggc3637la98yeas2809rpcw9"* ]]; then
    echo "$lcmd - sanity - 8 python - address (apple8,orange2,0,12345) address = OK"
  else
    echo $out
  fi

  out=$($lcmd address --password apple8 --second orange2 --iterations 12345 --mode python --script p2wpkh-p2sh)
  if [[ $out == *"3396HyGWCFA8iTYUdM9jFJnZRfGbSxXPA4"* ]]; then
    echo "$lcmd - sanity - 9 python - p2wpkh-p2sh address (apple8,orange2,0,12345) address = OK"
  fi

  out=$($lcmd address --password apple8 --second orange2 --iterations 12345 --mode python --script p2pkh)
  if [[ $out == *"1EXmNr6FutKG5ei1adFF5tnUwFaPk2Mcon"* ]]; then
    echo "$lcmd - sanity - 10 python - p2pkh address (apple8,orange2,0,12345) address = OK"
  fi

done
