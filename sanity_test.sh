#!/bin/bash
# Sanity tests to protect from breaking changes to the scripts

CMD="python3 -m twop"

out=$($CMD seed --password apple8)
if [[ $out == *"b930218bd7b1b6620ebc7bf0c225a76e5c3780411ea104cc83bd40d57f6d6f6d"* ]]; then
  echo "sanity - 1 - seed password (apple8) hex = OK"
fi
if [[ $out == *"rice library glass quarter bread country depend month valid bachelor hawk syrup sell scan afford stage age goat urban like fitness repeat sad poverty"* ]]; then
  echo "sanity - 2 - seed password (apple8) seed = OK"
fi

out=$($CMD seed --password apple8 --iterations 12345)
if [[ $out == *"df43c0019f4c8a04132779cc871b2645b0107f8452770ed4505976e0269efa03"* ]]; then
  echo "sanity - 3 - seed password (apple8,12345) hex = OK"
fi
if [[ $out == *"tent bulk about direct silly acoustic erode upset smart decide sister merge absurd divert bacon exclude attract penalty bind universe act exhaust trend improve"* ]]; then
  echo "sanity - 4 - seed password (apple8,12345) seed = OK"
fi

out=$($CMD address --password apple8 --second orange2)
if [[ $out == *"bc1q64zn323why9k4nutm63xscx82yufrnqmdun6xp"* ]]; then
  echo "sanity - 5 - address (apple8,orange2) address = OK"
fi

out=$($CMD address --password apple8 --second orange2 --num 0)
if [[ $out == *"bc1q64zn323why9k4nutm63xscx82yufrnqmdun6xp"* ]]; then
  echo "sanity - 6 - address (apple8,orange2,0) address = OK"
fi

out=$($CMD address --password apple8 --second orange2 --num 1)
if [[ $out == *"bc1qhgvpf9p97qs5vajdt3shy3hkssrfzh5e47u50z"* ]]; then
  echo "sanity - 7 - address (apple8,orange2,1) address = OK"
fi

out=$($CMD address --password apple8 --second orange2 --iterations 12345)
if [[ $out == *"bc1qmwd26qurpd2s8j693q3sct5pfkwlnkudu4np4p"* ]]; then
  echo "sanity - 8 - address (apple8,orange2,0,12345) address = OK"
fi
