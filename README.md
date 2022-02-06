# twopasswords

**WARNING Do Not Use this unless you understand the code !! I am not responsible for any funds you lose. WARNING**

A simple CLI to make secure Bitcoin brain seeds & wallets using two passwords.
- Make seed and wallet
- Show address(es) by number
- Supports Bitcoin
- Supports Legacy, P2SH Segwit & Segwit

_Why? How?_

"In some situations the safest way to carry Bitcoin around is with a password or two."

Using PBKDF2 to stretch the first password to make a random seed.  Then second password is then used as BIP39 passphrase.


## Usage

Make a seed from a password

```shell
./twopasswords seed --password="qwerty1"

./twopasswords seed --password="qwerty1" --iterations 123456
```

Make a HD wallet and show address _n_

```shell
./twopasswords address --password="qwerty1" --second="password" --num=0

./twopasswords address --password="qwerty1" --second="password" --num=0 --iterations 123456
```
