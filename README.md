# twopasswords: Simple brain wallet demo

![Build Status](https://github.com/timchurchard/twopasswords/workflows/Test/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-67.8%25-yellow)
[![License](https://img.shields.io/github/license/timchurchard/twopasswords)](/LICENSE)
[![Release](https://img.shields.io/github/release/timchurchard/twopasswords.svg)](https://github.com/timchurchard/twopasswords/releases/latest)
[![GitHub Releases Stats of twopasswords](https://img.shields.io/github/downloads/timchurchard/twopasswords/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=timchurchard&repository=twopasswords)


**WARNING Do Not Use this unless you understand the code !! I am not responsible for any funds you lose. WARNING**

This is a demo of a simple brain wallet system comprising a simple CLI to make Bitcoin brain wallets using two passwords.
- Make seed and wallet
- Show address(es) by number
- Supports Bitcoin
- Supports Legacy, P2SH Segwit & Segwit
- Optional BIP38 Encrypt WIF
- Wallet and utility functions

### What

One criticism of BIP39 is the source of entropy. A user may provide entropy from a weak source (password). This demo is an attempt to explore that option.

PBKDF2 is used to stretch the first password to make a 'random' seed. The stretching uses a large number of iterations and the user may specify that number.

Then second password is then used as BIP39 passphrase. In theory this combination makes brute forcing much harder.

Many people remember more than two passwords already and long numbers, so they should be able to safely store a long-term Bitcoin wallet in their brain too.

### Iterations

Iterations is the number of times to hash passwords. This increases the time & CPU effort required to brute force wallets made in this way. I suggest picking a large number (millions or billions) that you can remember. The larger the number the longer the calculation will take. Use 21,000,000 (million) for example, this takes about 10s for me. Using 21,000,000,000 (billion) takes 10,900 seconds (3 hours).

## Docker usage

```shell
$ docker build -t tc/twopasswords .

$ docker run -it --rm tc/twopasswords /examples/generate-addr.sh
Mnemonic = between emotion state blast corn question advice cement gesture future will wrong (Bip39 with second password: 112526)
Made address 0 (m/84'/0'/0'/0/0) = bc1qcpepzkt6hez3mw4lgcmp44scqmprpp8sdf2v6l

# Run the other examples
# docker run -it --rm tc/twopasswords /examples/notewallet/make_notewallet.sh
```

## Usage

Make a seed from a password

```shell
./twopasswords seed --password="qwerty1"

./twopasswords seed --password="qwerty1" --iterations 123456
```

Make an HD wallet and show address _n_

```shell
./twopasswords address --password="qwerty1" --second="password" --num=0

./twopasswords address --password="qwerty1" --second="password" --num=0 --iterations 123456
```

Decrypt bip38 key. Note it is not possible to know if the password is wrong for the decryption, you can specify the expected address to prove the password is right.
```shell
./twopasswords address --password="qwerty1" --second="password" --num=0 --iterations 123456 --bip38 password
#Mnemonic = scan note ramp aerobic insect cycle provide void nurse head couple pet sand favorite hedgehog educate melt illness verify fog denial tuition water enrich (Bip39 with second password: password)
#Made address 0 (m/84'/0'/0'/0/0) = bc1q3m7smsulgkc5tkxw3v82c2z5gll8c32qglfxdc
#WIF: 6PRNEXpCM9oAG1HhUffPPHeZYbaRJViw75inCmbdCPDPAkpiDcE8VvpSth

./twopasswords bip38 -b 6PRNEXpCM9oAG1HhUffPPHeZYbaRJViw75inCmbdCPDPAkpiDcE8VvpSth -p password -a bc1q3m7smsulgkc5tkxw3v82c2z5gll8c32qglfxdc
#Bitcoin P2PKH:                  5KdoEi385k3ACP492eyGYhUMvhiyEh9bPvd4MGGZUZm3i6GtSAE
#Bitcoin P2PKH (Compressed):     L5FR5W8NFvxXbELrSJbMcudmN2kFDSCvpBg9nSPgLfbQx7DfzA59
#Bitcoin P2WPKH:                 p2wpkh:L5FR5W8NFvxXbELrSJbMcudmN2kFDSCvpBg9nSPgLfbQx7DfzA59
```

## Examples

Some example scripts are provided to explore what can be done with this utility

### notewallet

[./examples/notewallet](./examples/notewallet) creates a random number of iterations and two random passwords then lists several addresses. This is an example of a wallet in a text file where the secrets could be easily written on a paper note.

### paperwallet

[./examples/paperwallet](./examples/paperwallet) create example paperwallet (single public & private key)

## Utility functions

### balance

Attempt to get the balance of a wallet

```shell
./twopasswords balance --password="qwerty1" --second="password" --iterations 123456
```

