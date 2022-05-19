# Copyright Tim Churchard 2020

import os
from collections import namedtuple

electrum = None
try:
    import electrum
    from electrum import keystore
    from electrum.storage import WalletStorage
    from electrum.wallet_db import WalletDB
    from electrum.wallet import Wallet
    from electrum.simple_config import SimpleConfig
except ImportError:
    pass

from cryptotools.BTC import Xprv
from mnemonic import Mnemonic

from .const import MIN_LEN_PASSWORD


AddressResult = namedtuple('AddressResult', 'address wif num')


def make_address_py(mnemonic: str, second: str, num: int = 0, script: str = 'p2wpkh', hardened: bool = False):
    script = script.upper()
    m = Xprv.from_mnemonic(mnemonic, passphrase=second, addresstype=script)

    if hardened:
        num = float(num)

    if script == 'P2WPKH':
        addr = (m/84./0./0./0/num)
    elif script == 'P2WPKH-P2SH':
        addr = (m/49./0./0./0/num)
    elif script == 'P2PKH':
        addr = (m/44./0./0./0/num)

    return AddressResult(address=addr.address(script), wif=f'{script}:{addr.key.wif(compressed=True)}', num=num)


def make_address_electrum(mnemonic: str, second: str, num: int = 0, script: str = 'p2wpkh', path: str = 'wallet.db', remove: bool = True):
    if electrum is None:
        raise ImportError('Unable to import electrum.  Follow README instructions.')

    if len(second) < MIN_LEN_PASSWORD:
        raise ValueError(f'Password too short {len(second)} < {MIN_LEN_PASSWORD}')
    script_types = ('p2pkh', 'p2wpkh', 'p2wpkh-p2sh')
    if script not in script_types:
        raise ValueError(f'Script {script} must be noe of {script_types}')

    if os.path.exists(path):
        os.unlink(path)

    language = 'english'
    seed = Mnemonic(language).to_seed(mnemonic).hex()

    derivation = None
    if script == 'p2wpkh':
        derivation = "m/84'/0'/0'"
    elif script == 'p2wpkh-p2sh':
        derivation = "m/49'/0'/0'"
    elif script == 'p2pkh':
        script = 'standard'
        derivation = "m/44'/0'/0'"

    ks = keystore.from_bip39_seed(mnemonic, second, derivation, xtype=script)

    db = WalletDB('', manual_upgrades=False)
    db.put('keystore', ks.dump())
    db.put('wallet_type', 'standard')

    storage = WalletStorage(path)
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
