# Copyright Tim Churchard 2020

from collections import namedtuple
from hashlib import sha256, pbkdf2_hmac
from time import monotonic, sleep
from threading import Thread, Event

from mnemonic import Mnemonic

from .const import DEF_ITS_SALT, DEF_ITS_PBKDF2, MIN_LEN_PASSWORD, MIN_ITS_PBKDF2, MIN_ITS_SALT, DEF_VERBOSE_TIME


def make_salt(password: bytes, iterations: int = DEF_ITS_SALT) -> (bytes, int):
    """make_salt: Takes password bytes and hashes iterations times to produce a salt

    sha256(password) -> sha256(result) -> sha256(result) until iterations

    Returns (salt bytes, seconds (zero if less than DEF_VERBOSE_TIME))
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
            print(f'(1/2) {progress:.2f}% SALT')

    total_secs = last_time - start_time

    return result, total_secs


SeedResult = namedtuple('SeedResult', 'bytes hex mnemonic')


def make_seed(password: bytes,
              salt: bytes,
              iterations: int = DEF_ITS_PBKDF2,
              language: str = 'english',
              expected_secs: int = 0) -> SeedResult:
    """make_seed: Takes password and salt and returns SeedResult namedtuple

    if expected_secs > 0 then show a timer as progress percent
    """
    if len(password) < MIN_LEN_PASSWORD:
        raise ValueError(f'Password too short {len(password)} < {MIN_LEN_PASSWORD}')
    if iterations < MIN_ITS_PBKDF2:
        raise ValueError(f'Seed iterations too low {iterations} < {MIN_ITS_PBKDF2}')
    # todo: Make this configurable?
    if len(salt) != 32:
        raise ValueError('Salt must be 32-bytes')

    # Progress indicator thread
    thread = None
    evt = Event()
    if expected_secs:

        # Magic! Reduce the expected_secs because pbkdf2 is faster than the make_salt loop.
        expected_secs = expected_secs * 0.75

        def show_progress(expected_secs, evt):
            start_time = last_time = monotonic()
            while not evt.is_set():
                if monotonic() > last_time + DEF_VERBOSE_TIME:
                    last_time = monotonic()
                    progress = (100 / expected_secs) * (last_time - start_time)
                    print(f'(2/2) {progress:.2f}% PBKDF2')
                sleep(9.5)
        thread = Thread(target=show_progress, args=(expected_secs, evt,))
        thread.start()

    dk = pbkdf2_hmac('sha256', password, salt, iterations)
    seed = Mnemonic(language).to_mnemonic(dk)

    evt.set()
    if thread is not None:
        thread.join()

    return SeedResult(bytes=dk, hex=dk.hex(), mnemonic=seed)
