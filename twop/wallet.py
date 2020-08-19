# Copyright Tim Churchard 2020

import os
from collections import namedtuple

from electrum import keystore
from electrum.storage import WalletStorage
from electrum.wallet_db import WalletDB
from electrum.wallet import Wallet
from electrum.simple_config import SimpleConfig
from mnemonic import Mnemonic

from .const import MIN_LEN_PASSWORD


AddressResult = namedtuple('AddressResult', 'address wif num')


def make_address(mnemonic: str, second: str, num: int = 0, script: str = 'p2wpkh', path: str = 'wallet.db', remove: bool = True):
    if len(second) < MIN_LEN_PASSWORD:
        raise ValueError(f'Password too short {len(second)} < {MIN_LEN_PASSWORD}')
    script_types = ('p2pkh', 'p2wpkh', 'p2wpkh-p2sh')
    if script not in script_types:
        raise ValueError(f'Script {script} must be noe of {script_types}')

    # TODO: Right now we have to delete the wallet path, if it exists
    # - Does not seem to be encrypted? But all the password related funcs raise invalid padding ??
    if os.path.exists(path):
        os.unlink(path)

    language = 'english'  # todo
    seed = Mnemonic(language).to_seed(mnemonic).hex()

    storage = WalletStorage(path)
    db = WalletDB('', manual_upgrades=False)
    ks = keystore.from_bip39_seed(mnemonic, second, "m/84'/0'/0'/0")

    db.put('keystore', ks.dump())
    db.put('wallet_type', 'standard')

    wallet = Wallet(db, storage, config=SimpleConfig())
    if not storage.file_exists():
        wallet.update_password(old_pw=None, new_pw=second, encrypt_storage=True)
    wallet.synchronize()
    wallet.save_db()

    addr = wallet.get_receiving_addresses()[num]
    wif = wallet.export_private_key(addr, password=second)

    if remove:
        os.unlink(path)

    return AddressResult(address=addr, wif=wif, num=num)
