package main

import (
	"flag"
	"fmt"
	"github.com/timchurchard/twopasswords/cmd"
	"os"
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
	}

	usageRoot()
}

func usageRoot() {
	fmt.Printf("usage: %s command(seed|wallet) options\n", cliName)
	os.Exit(1)
}
