# Copyright Tim Churchard 2020

from .const import ERROR, OK, DEF_ITS_PBKDF2
from .seed import make_salt, make_seed
from .wallet import make_address

import click


@click.group()
def main():
    pass



@click.command()
@click.option('--password', help='Password for seed')
@click.option('--iterations', default=DEF_ITS_PBKDF2, help='Number of iterations for PBKDF2')
def seed(password: str, iterations: int):
    password_bytes = password.encode('utf8')

    try:
        salt, expected_secs = make_salt(password_bytes, iterations=iterations)
    except ValueError as exc:
        print(exc)
        return ERROR

    try:
        seed = make_seed(password_bytes, salt, iterations=iterations, expected_secs=expected_secs)
    except ValueError as exc:
        print(exc)
        return ERROR

    print(f'Made seed. Hex = {seed.hex}\n{seed.mnemonic}')
    return OK


@click.command()
@click.option('--password', help='Password for seed')
@click.option('--iterations', default=DEF_ITS_PBKDF2, help='Number of iterations for PBKDF2')
@click.option('--second', help='Password for wallet')
@click.option('--num', default=0, help='Address number')
@click.option('--script', default='p2wpkh', help='Script (p2pkh, p2wpkh, p2wpkh-p2sh)')
@click.option('--path', default='wallet.db', help='Path to wallet file')
@click.option('--rm', default=True, help='Remove the electrum wallet file')
@click.option('--verbose', default=False, help='Verbose mode shows the seed & WIF')
def address(password, second, iterations, num, script, path, rm, verbose):
    password_bytes = password.encode('utf8')

    try:
        salt, expected_secs = make_salt(password_bytes, iterations=iterations)
    except ValueError as exc:
        print(exc)
        return ERROR

    try:
        seed = make_seed(password_bytes, salt, iterations=iterations, expected_secs=expected_secs)
    except ValueError as exc:
        print(exc)
        return ERROR
    if verbose:
        print(f'Made seed. Hex = {seed.hex}\n{seed.mnemonic}\n')

    addr = make_address(seed.mnemonic, second, num, script, path, rm)

    print(f'Made address: {addr.num} = {addr.address}')
    if verbose:
        print(f'WIF: {addr.wif}')

    return OK


main.add_command(seed)
main.add_command(address)


if __name__ == '__main__':
    from sys import exit
    exit(main())
