package pkg

import (
	"errors"
)

// AddressResult holds the MakeAddress result
type AddressResult struct {
	Address        string
	SecretExponent []byte
	Wif            string
	Num            int
	DerivationPath string
}

// MakeAddress function takes mnemonic and password and returns an address and wif
func MakeAddress(mnemonic string, password string, num int, script string) (AddressResult, error) {
	const (
		// Seed is always 256-bit here
		bitSize = 256

		// Only produce compressed addresses
		compress = true

		// Only produce bitcoin addresses
		coinType = CoinTypeBTC
	)

	var (
		err     error
		purpose Purpose
		address string
		wif     string
	)

	km, err := NewKeyManager(bitSize, password, mnemonic)
	if err != nil {
		return AddressResult{}, err
	}

	switch script {
	case "p2wpkh":
		purpose = PurposeBIP84
	case "p2wpkh-p2sh":
		purpose = PurposeBIP49
	case "p2pkh":
		purpose = PurposeBIP44
	default:
		return AddressResult{}, errors.New("invalid script must be in ('p2wpkh', 'p2wpkh-p2sh', 'p2pkh')")
	}

	key, err := km.GetKey(purpose, coinType, 0, 0, uint32(num))
	if err != nil {
		return AddressResult{}, err
	}

	switch script {
	case "p2wpkh":
		wif, _, address, _, err = key.Encode(compress)
	case "p2wpkh-p2sh":
		wif, _, _, address, err = key.Encode(compress)
	case "p2pkh":
		wif, address, _, _, err = key.Encode(compress)
	}
	if err != nil {
		return AddressResult{}, err
	}

	return AddressResult{
		Address:        address,
		SecretExponent: key.bip32Key.Key,
		Wif:            wif,
		Num:            num,
		DerivationPath: key.GetPath(),
	}, nil
}
