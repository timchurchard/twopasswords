package cmd

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TestBip38Main(t *testing.T) {
	const cliName = "bip38"

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
				"-b",
				"6PRNEXpCM9oAG1HhUffPPHeZYbaRJViw75inCmbdCPDPAkpiDcE8VvpSth",
				"-p",
				"password",
				"-a",
				"bc1q3m7smsulgkc5tkxw3v82c2z5gll8c32qglfxdc",
			},
			want:    0,
			wantOut: "Bitcoin P2PKH:\t\t\t5KdoEi385k3ACP492eyGYhUMvhiyEh9bPvd4MGGZUZm3i6GtSAE\nBitcoin P2PKH (Compressed):\tL5FR5W8NFvxXbELrSJbMcudmN2kFDSCvpBg9nSPgLfbQx7DfzA59\nBitcoin P2WPKH:\t\t\tp2wpkh:L5FR5W8NFvxXbELrSJbMcudmN2kFDSCvpBg9nSPgLfbQx7DfzA59\n",
		},
		{
			// Secret exponent 009daf84b9b861a2a6f1de543781ebf2c0693707bb0dd9c55aff2931a754583a
			// begins with a zero byte. None of the BIP38 test vectors do, so
			// nothing else in the suite covers a 31-byte big.Int.
			name: "leading zero byte",
			args: []string{
				"-b",
				"6PRVhMBUXTbmLZsbsRXmDEoNmsEH6UQRSsMNbeaVn1C4b1TgUcMCReaDzx",
				"-p",
				"TestingOneTwoThree",
				"-a",
				"12HPHgPkaSDqZ93bEH4R7NHFrAqUeoppj6",
			},
			want:    0,
			wantOut: "Bitcoin P2PKH:\t\t\t5HpZKF9UfyNG9EpVaic45Af4GJeZ9LGuGbxEnxPU3zvys2NieiN\nBitcoin P2PKH (Compressed):\tKwEudXSR91nqqMzd9pdiZPcyQYtksn5BhsGosqUAQEhFyRc8AgiG\nBitcoin P2WPKH:\t\t\tp2wpkh:KwEudXSR91nqqMzd9pdiZPcyQYtksn5BhsGosqUAQEhFyRc8AgiG\n",
		},
	}
	for _, tt := range tests {
		// reset flags else panic
		flag.CommandLine = flag.NewFlagSet(cliName, flag.ExitOnError)
		os.Args = append([]string{cliName}, tt.args...)

		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if got := Bip38Main(out); got != tt.want {
				t.Errorf("Bip38Main() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Bip38Main() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
