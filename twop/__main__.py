# Copyright Tim Churchard 2020

from .const import ERROR, OK
from .seed import make_salt, make_seed
from .wallet import make_address

import click


@click.group()
def main():
    pass



@click.command()
@click.option('--password', help='Password for seed')
def seed(password: str):
    password_bytes = password.encode('utf8')

    try:
        salt = make_salt(password_bytes)
    except ValueError as exc:
        print(exc)
        return ERROR

    try:
        seed = make_seed(password_bytes, salt)
    except ValueError as exc:
        print(exc)
        return ERROR

    print(f'Made seed. Hex = {seed.hex}\n{seed.mnemonic}')
    return OK


@click.command()
@click.option('--password', help='Password for seed')
@click.option('--second', help='Password for wallet')
@click.option('--num', default=0, help='Address number')
@click.option('--script', default='p2wpkh', help='Script (p2pkh, p2wpkh, p2wpkh-p2sh)')
@click.option('--path', default='wallet.db', help='Path to wallet file')
@click.option('--rm', default=True, help='Remove the electrum wallet file')
def address(password, second, num, script, path, rm):
    password_bytes = password.encode('utf8')
    second_bytes = second.encode('utf8')

    try:
        salt = make_salt(password_bytes)
    except ValueError as exc:
        print(exc)
        return ERROR

    try:
        seed = make_seed(password_bytes, salt)
    except ValueError as exc:
        print(exc)
        return ERROR
    # print(f'Made seed. Hex = {seed.hex}\n{seed.mnemonic}\n')

    remove = rm.lower() == 'true'
    addr = make_address(seed.mnemonic, second, num, script, path, remove)

    print(f'Made address: {addr.num} = {addr.address}')
    # print(f'WIF: {addr.wif}')

    return OK


main.add_command(seed)
main.add_command(address)


if __name__ == '__main__':
    from sys import exit
    exit(main())
