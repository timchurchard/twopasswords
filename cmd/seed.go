package cmd

import (
	"flag"
	"fmt"
	"io"

	"github.com/timchurchard/twopasswords/pkg"
)

func SeedMain(out io.Writer) int {
	var (
		iterations int
		bits       int
		password   string
	)

	flag.IntVar(&iterations, "iterations", pkg.DefaultIterations, usageIterations)
	flag.IntVar(&iterations, "i", pkg.DefaultIterations, usageIterations+" (shorthand)")

	flag.StringVar(&password, "password", defaultEmpty, usagePassword)
	flag.StringVar(&password, "p", defaultEmpty, usagePassword+" (shorthand)")

	flag.IntVar(&bits, "bits", default256, usageBits)
	flag.IntVar(&bits, "b", default256, usageBits+" (shorthand)")

	flag.Parse()

	seedResult, err := makeSeed(out, password, iterations, bits)
	if err != nil {
		return 1
	}

	fmt.Fprintf(out, "Made seed. Hex = %x\n", seedResult.Entropy)
	fmt.Fprintf(out, "Mnemonic = %s\n", seedResult.Mnemonic)

	return 0
}

func makeSeed(out io.Writer, password string, iterations, bits int) (pkg.SeedResult, error) {
	passwordBytes := []byte(password)
	saltBytes, err := pkg.MakeSalt(passwordBytes, iterations)
	if err != nil {
		fmt.Fprintf(out, "Error making salt: %v", err)
		return pkg.SeedResult{}, err
	}

	seedResult, err := pkg.MakeSeed([]byte(password), saltBytes, iterations, bits, "english")
	if err != nil {
		fmt.Fprintf(out, "Error making seed: %v", err)
		return pkg.SeedResult{}, err
	}

	return seedResult, nil
}
