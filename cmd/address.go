package cmd

import (
	"flag"
	"fmt"
	"github.com/timchurchard/twopasswords/pkg"
	"io"
)

func AddressMain(out io.Writer) int {
	const (
		defaultEmpty    = ""
		usageIterations = "Number of iterations for PBKDF2"
		usagePassword   = "Password for seed"
		usageSecond     = "Password for bip39 and wallet encryption"
		usageNum        = "Address number to make"
		usageScript     = "Script (p2pkh, p2wpkh, p2wpkh-p2sh)"
		defaultScript   = "p2wpkh"
		usagePath       = "Path to wallet file"
		usageRm         = "Remove the electrum wallet file"
		usageMode       = "Mode is not used in golang impl. Only for compatibility with python CLI."
	)

	var (
		verbose    bool
		iterations int
		password   string
		second     string
		num        int
		script     string
		path       string
		remove     bool
		mode       string
	)

	flag.BoolVar(&verbose, "v", false, "Verbose mode")

	flag.IntVar(&iterations, "iterations", pkg.DefaultIterations, usageIterations)
	flag.IntVar(&iterations, "i", pkg.DefaultIterations, usageIterations+" (shorthand)")

	flag.StringVar(&password, "password", defaultEmpty, usagePassword)
	flag.StringVar(&password, "p", defaultEmpty, usagePassword+" (shorthand)")

	flag.StringVar(&second, "second", defaultEmpty, usageSecond)
	flag.StringVar(&second, "s", defaultEmpty, usageSecond+" (shorthand)")

	flag.IntVar(&num, "num", 0, usageNum)
	flag.IntVar(&num, "n", 0, usageNum+" (shorthand)")

	flag.StringVar(&mode, "mode", defaultEmpty, usageMode)
	flag.StringVar(&mode, "m", defaultEmpty, usageMode+" (shorthand)")

	flag.StringVar(&script, "script", defaultScript, usageScript)
	flag.StringVar(&path, "path", defaultEmpty, usagePath)
	flag.BoolVar(&remove, "rm", true, usageRm)

	flag.Parse()

	seedResult, err := makeSeed(out, password, iterations)
	if err != nil {
		return 1
	}

	if verbose {
		fmt.Fprintf(out, "Made seed. Hex = %x\n", seedResult.Entropy)
	}

	addressResult, err := pkg.MakeAddress(seedResult.Mnemonic, second, num, script)
	if err != nil {
		fmt.Fprintf(out, "Error making address: %v", err)
		return 1
	}

	fmt.Fprintf(out, "Mnemonic = %s (Bip39 with second password: %s)\n", seedResult.Mnemonic, second)
	fmt.Fprintf(out, "Made address %d (%s) = %s\n", addressResult.Num, addressResult.DerivationPath, addressResult.Address)
	fmt.Fprintf(out, "WIF: %s\n", addressResult.Wif)

	return 0
}
