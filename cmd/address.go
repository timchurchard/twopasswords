package cmd

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"

	bip38address "github.com/sour-is/bitcoin/address"
	"github.com/sour-is/bitcoin/bip38"
	"github.com/timchurchard/twopasswords/pkg"
)

type addressFlagData struct {
	Verbose    bool
	Bits       int
	Iterations int
	Password   string
	Second     string
	Num        int
	Script     string
	Encrypt    string
}

func addressFlags() addressFlagData {
	result := addressFlagData{}

	flag.BoolVar(&result.Verbose, "v", false, "Verbose mode")

	flag.IntVar(&result.Iterations, "iterations", pkg.DefaultIterations, usageIterations)
	flag.IntVar(&result.Iterations, "i", pkg.DefaultIterations, usageIterations+" (shorthand)")

	flag.IntVar(&result.Bits, "bits", default256, usageBits)
	flag.IntVar(&result.Bits, "b", default256, usageBits+" (shorthand)")

	flag.StringVar(&result.Password, "password", defaultEmpty, usagePassword)
	flag.StringVar(&result.Password, "p", defaultEmpty, usagePassword+" (shorthand)")

	flag.StringVar(&result.Second, "second", defaultEmpty, usageSecond)
	flag.StringVar(&result.Second, "s", defaultEmpty, usageSecond+" (shorthand)")

	flag.IntVar(&result.Num, "num", 0, usageNum)
	flag.IntVar(&result.Num, "n", 0, usageNum+" (shorthand)")

	flag.StringVar(&result.Script, "script", defaultScript, usageScript)

	flag.StringVar(&result.Encrypt, "bip38", defaultEmpty, usageBip38)

	mode := "unused"
	flag.StringVar(&mode, "mode", defaultEmpty, "unused. Left for compatibility.")

	flag.Parse()

	return result
}

func AddressMain(out io.Writer) int {
	const minBip38Len = 4 // Minimum length of bip38 password

	args := addressFlags()

	seedResult, err := makeSeed(out, args.Password, args.Iterations, args.Bits)
	if err != nil {
		return 1
	}

	if args.Verbose {
		fmt.Fprintf(out, "Made seed. Hex = %x\n", seedResult.Entropy)
	}

	addressResult, err := pkg.MakeAddress(seedResult.Mnemonic, args.Second, args.Num, args.Script)
	if err != nil {
		fmt.Fprintf(out, "Error making address: %v", err)
		return 1
	}

	fmt.Fprintf(out, "Mnemonic = %s (Bip39 with second password: %s)\n", seedResult.Mnemonic, args.Second)
	fmt.Fprintf(out, "Made address %d (%s) = %s\n", addressResult.Num, addressResult.DerivationPath, addressResult.Address)

	if args.Encrypt == "" {
		fmt.Fprintf(out, "WIF: %s\n", addressResult.Wif)
	} else if len(args.Encrypt) < minBip38Len {
		fmt.Fprintf(out, "Error bip38 password too short %d < %d", len(args.Encrypt), minBip38Len)
		return 1
	} else {
		priv, err := bip38address.ReadPrivateKey(hex.EncodeToString(addressResult.SecretExponent))
		if err != nil {
			fmt.Fprintf(out, "Error decoding WIF for bip38 %v", err)
			return 1
		}

		fmt.Fprintf(out, "WIF: %s\n", bip38.Encrypt(priv, args.Encrypt))
	}

	return 0
}
