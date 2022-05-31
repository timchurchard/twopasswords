#!/usr/bin/env bash

cmd="../../twopasswords"

template_img="./bitcoin_piggy_1_t2_export.png"

font_reqular="FreeMono"
font_bold="FreeMono-Bold"

tag=$((($RANDOM%9000)+1000))
finalfile="output_$tag.png"
if test -f "$finalfile"; then
  exit
fi

iterations="420699"

first=$(tr -cd '[:alnum:]' < /dev/urandom | fold -w30 | head -n1)
second=$(tr -cd '[:alnum:]' < /dev/urandom | fold -w30 | head -n1)
third=$(tr -cd 'A-Z0-9' < /dev/urandom | fold -w10 | head -n1)

tmpfile=$(mktemp /tmp/paper-script.XXXXXX)
$cmd address --script p2wpkh --iterations $iterations --password $first --second $second --bip38 $third > $tmpfile

address=$(cat $tmpfile | grep bc1q | awk '{print $6}')
wif=$(cat $tmpfile | grep WIF | awk '{print $2}')
rm $tmpfile

# Note: 111 x 111 px default is small, we'll double
public=$(mktemp /tmp/paper-script.XXXXXX)
qrencode -o $public $address
convert $public -resize 222x222 $public

# Note: 123 x 123 px default
private=$(mktemp /tmp/paper-script.XXXXXX)
qrencode -o $private $wif
convert $private -resize 246x246 $private

# Piggy Template
$(convert $template_img \
  \( "$public" \) -geometry +125+1375 -composite \
  \( "$private" \) -geometry +1225+1500 -composite \
  $finalfile )
rm "$public" "$private"

$(convert $finalfile \
    -font "$font_reqular" -fill black -pointsize 28 -draw "text 125,125 '$tag'" $finalfile)

$(convert $finalfile \
    -font "$font_reqular" -fill black -pointsize 32 -draw "text 350,1500 '$address'" $finalfile)

$(convert $finalfile \
    -font "$font_bold" -fill black -pointsize 32 -draw "text 125,1760 '$wif'" $finalfile)

echo "Made $tag = $third ($address $wif) $finalfile"