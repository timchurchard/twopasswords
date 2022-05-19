# twopasswords

**WARNING Do Not Use this unless you understand the code !! I am not responsible for any funds you lose. WARNING**

A simple CLI to make secure Bitcoin brain seeds & wallets using two passwords.
- Make seed and wallet
- Show address(es) by number
- Supports Bitcoin
- Supports Legacy, P2SH Segwit & Segwit
- Optional BIP38 Encrypt WIF

_Why? How?_

"In some situations the safest way to carry Bitcoin around is with a password or two."

Using PBKDF2 to stretch the first password to make a random seed.  Then second password is then used as BIP39 passphrase.


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