# twopasswords

A simple CLI and python module to make a secure Bitcoin wallet from two passwords.
- Make seed and wallet
- Show address(es) by number
- Supports Bitcoin, Litecoin, Dogecoin
- Supports Legacy, P2SH Segwit & Segwit

*WARNING* Do Not Use this unless you understand the code !! I am not responsible for any funds you lose. *WARNING*

_Why? How?_

In some situations the safest way to carry Bitcoin around is with a password or two.

Using PBKDF2 to stretch the first password to make a random seed.  Then second password as BIP39 password.


## Hacking

Fetch electrum git submodule

```shell
git submodule update --init
```

Create a venv and install requirements.

```shell
python3.6 -m venv venv
. venv37/bin/activate
python3 -m pip install -U pip setuptools
pip install -r requirements.txt
```

Install electrum

```shell
cd 3rd/electrum
pip install -e .
```

Create an alias

```shell
alias twop="python3 -m twop"
```

## Usage

twop seed --password="qwerty1"

twop address --password="qwerty1" --second="password" --num=0
