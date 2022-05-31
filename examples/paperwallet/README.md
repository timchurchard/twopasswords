# examples/paperwallet: Simple paper wallet generator example

A simple example using [twopasswords](https://github.com/timchurchard/twopasswords) to make keys for a paperwallet. And use imagemagick and qrencode to render the final image on top of a template image.

## Setup and Usage

```shell
# Imagemagick and qrencode are required:
sudo apt install imagemagick qrencode

# Notice that ../../twopasswords will be used by default
./make_paperwallet.sh
```

The script will produce a line of output like this
```shell
Made 6947 = ETGX617C4I (bc1qjwrfaz5p6gqltmempzl2tys7esjfwugy983qjg 6PRWNUTZTQRw6D3DZZ9yzHWN7uSFsqpVBhGVFS8kBsJM36MBCwztxfzQbL) output_6947.png`
     ^ tag  ^ BIP38 password  ^ Address                             ^ BIP38 encrypted private key 
```

## Decrypt BIP38

```shell
./twopasswords bip38 -b 6PRWNUTZTQRw6D3DZZ9yzHWN7uSFsqpVBhGVFS8kBsJM36MBCwztxfzQbL -a bc1qjwrfaz5p6gqltmempzl2tys7esjfwugy983qjg -p ETGX617C4I
```

## Hacking

This is an example. There are lots of hardcoded values and assumptions in the make_password.sh script. I have tried to make it clear and easy to understand. If you wish to change the template image you may also have to change the text location, font sizes, etc.