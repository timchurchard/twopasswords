package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/timchurchard/twopasswords/cmd"
)

const cliName = "twopasswords"

func main() {
	if len(os.Args) < 2 {
		usageRoot()
	}

	// Save the command and reset the flags
	command := os.Args[1]
	flag.CommandLine = flag.NewFlagSet(cliName, flag.ExitOnError)
	os.Args = append([]string{cliName}, os.Args[2:]...)

	switch command {
	case "seed":
		os.Exit(cmd.SeedMain(os.Stdout))
	case "address":
		os.Exit(cmd.AddressMain(os.Stdout))
	case "bip38":
		os.Exit(cmd.Bip38Main(os.Stdout))
	case "balance":
		os.Exit(cmd.BalanceMain(os.Stdout))
	}

	usageRoot()
}

func usageRoot() {
	fmt.Printf("usage: %s command(seed|address|bip38|wallet) options\n", cliName)
	os.Exit(1)
}
