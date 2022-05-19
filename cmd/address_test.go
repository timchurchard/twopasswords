package cmd

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TestAddressMain(t *testing.T) {
	const cliName = "address"

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	tests := []struct {
		name    string
		args    []string
		want    int
		wantOut string
	}{
		{
			name: "sanity",
			args: []string{
				"--password",
				"qwerty1",
				"--second",
				"second",
			},
			want:    0,
			wantOut: "Mnemonic = despair able drum rug inside quarter daring history remain youth kangaroo spend chest rib hover winner mixed payment tragic little option treat toilet attract (Bip39 with second password: second)\nMade address 0 (m/84'/0'/0'/0/0) = bc1qpdr30gfqyqlg44jwl4avqjejtqdc85c5hkcygn\nWIF: KwPBW3pQBG8JHRGhqyuUrQVRidc2j7NzsiJi7zH2qtVqnPjAtANZ\n",
		},
	}
	for _, tt := range tests {
		// reset flags else panic
		flag.CommandLine = flag.NewFlagSet(cliName, flag.ExitOnError)
		os.Args = append([]string{cliName}, tt.args...)

		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if got := AddressMain(out); got != tt.want {
				t.Errorf("AddressMain() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("AddressMain() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
