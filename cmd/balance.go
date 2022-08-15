package cmd

import (
	"fmt"
	"io"

	opendime "github.com/timchurchard/opendime-utils/pkg"
	"github.com/timchurchard/twopasswords/pkg"
)

func BalanceMain(out io.Writer) int {
	const (
		fiat = "GBP" // TODO
		gap  = 20
	)

	args := addressFlags()

	seedResult, err := makeSeed(out, args.Password, args.Iterations, args.Bits)
	if err != nil {
		return 1
	}

	totalBitcoin := 0.0
	totalFiat := 0.0

	lastFound := 0
	idx := 0

	for {
		addressResult, err := pkg.MakeAddress(seedResult.Mnemonic, args.Second, idx, args.Script)
		if err != nil {
			fmt.Fprintf(out, "Error making address: %v", err)
			return 1
		}

		bitcoinAmount, fiatAmount, _, _ := opendime.CheckBalance(addressResult.Address, fiat)
		if bitcoinAmount > 0 {
			lastFound = idx
		}

		totalBitcoin += bitcoinAmount
		totalFiat += fiatAmount

		fmt.Printf("%02d | %s | %.8f | %.2f %s \n", idx, addressResult.Address, bitcoinAmount, fiatAmount, fiat)

		idx += 1
		if idx > lastFound+gap {
			break
		}
	}

	fmt.Printf("Total: %.8f %.2f\n", totalBitcoin, totalFiat)

	return 0
}
