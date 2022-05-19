# Copyright Tim Churchard 2020

from collections import namedtuple
from hashlib import sha256, pbkdf2_hmac
from time import monotonic

from mnemonic import Mnemonic

from .const import DEF_ITS_SALT, DEF_ITS_PBKDF2, MIN_LEN_PASSWORD, MIN_ITS_PBKDF2, MIN_ITS_SALT, DEF_VERBOSE_TIME


def make_salt(password: bytes, iterations: int = DEF_ITS_SALT) -> bytes:
    """make_salt: Takes password bytes and hashes iterations times to produce a salt

    sha256(password) -> sha256(result) -> sha256(result) until iterations
    """
    if len(password) < MIN_LEN_PASSWORD:
        raise ValueError(f'Password too short {len(password)} < {MIN_LEN_PASSWORD}')
    if iterations < MIN_ITS_SALT:
        raise ValueError(f'Salt iterations too low {iterations} < {MIN_ITS_SALT}')

    start_time = last_time = monotonic()
    result = password
    for i in range(iterations):
        md = sha256()
        md.update(result)
        result = md.digest()

        if monotonic() > last_time + DEF_VERBOSE_TIME:
            last_time = monotonic()
            progress = (100 / iterations) * i
            print(f'(1/2) {progress:.2f}% SALT {i} / {iterations}')

    if not start_time == last_time:
        print(f'SALT total seconds {last_time - start_time}')

    return result


SeedResult = namedtuple('SeedResult', 'bytes hex mnemonic')


def make_seed(password: bytes, salt: bytes, iterations: int = DEF_ITS_PBKDF2, language: str = 'english') -> SeedResult:
    """make_seed: Takes password and salt and returns namedtuple
    """
    if len(password) < MIN_LEN_PASSWORD:
        raise ValueError(f'Password too short {len(password)} < {MIN_LEN_PASSWORD}')
    if iterations < MIN_ITS_PBKDF2:
        raise ValueError(f'Seed iterations too low {iterations} < {MIN_ITS_PBKDF2}')
    # todo: Make this configurable?
    if len(salt) != 32:
        raise ValueError(f'Salt must be 32-bytes')

    # TODO: If the make_salt timer happened it would be nice to have a timer here too
    dk = pbkdf2_hmac('sha256', password, salt, iterations)
    seed = Mnemonic(language).to_mnemonic(dk)

    return SeedResult(bytes=dk, hex=dk.hex(), mnemonic=seed)
