package cmd

import (
	"flag"
	"fmt"
	"io"

	"github.com/sour-is/bitcoin/bip38"
	opendime "github.com/timchurchard/opendime-utils/pkg"
)

func Bip38Main(out io.Writer) int {
	const (
		defaultEmpty  = ""
		usageBip38Wif = "BIP38 Encrypted WIF"
		usagePassword = "BIP38 Password"
		usageAddress  = "Expected address"

		prefixBitcoinHex = "80"
	)

	var (
		bip38wif string
		password string
		address  string
	)

	flag.StringVar(&bip38wif, "bip38wif", defaultEmpty, usageBip38Wif)
	flag.StringVar(&bip38wif, "b", defaultEmpty, usageBip38Wif+" (shorthand)")

	flag.StringVar(&password, "password", defaultEmpty, usagePassword)
	flag.StringVar(&password, "p", defaultEmpty, usagePassword+" (shorthand)")

	flag.StringVar(&address, "address", defaultEmpty, usageAddress)
	flag.StringVar(&address, "a", defaultEmpty, usageAddress+" (shorthand)")

	flag.Parse()

	priv, err := bip38.Decrypt(bip38wif, password)
	if err != nil {
		fmt.Fprintf(out, "Error on bip38.Decrypt: %v", err)
		return 1
	}

	_, secretExponentHex, _, err := opendime.ValidateWif(priv.String())
	if err != nil {
		fmt.Fprintf(out, "Error reading WIF: %v", err)
		return 1
	}

	p2pkh := opendime.ToWif(prefixBitcoinHex, secretExponentHex, false)
	p2pkhComp := opendime.ToWif(prefixBitcoinHex, secretExponentHex, true)

	if address != "" {
		addresses, _ := opendime.GetAddresses(opendime.VerifiedMessage{
			PublicKeyHex: priv.PublicKey.String(), // todo: Bit hacky. GetAddresses only uses this field currently.
		})
		if address != addresses.BitcoinP2PKH &&
			address != addresses.BitcoinP2PKHCompressed &&
			address != addresses.BitcoinP2WPKH {
			fmt.Fprintf(out, "Error expected address not found. Probably wrong password.")
			return 1
		}
	}

	fmt.Fprintf(out, "Bitcoin P2PKH:\t\t\t%s\n", p2pkh)
	fmt.Fprintf(out, "Bitcoin P2PKH (Compressed):\t%s\n", p2pkhComp)
	fmt.Fprintf(out, "Bitcoin P2WPKH:\t\t\tp2wpkh:%s\n", p2pkhComp)

	return 0
}
