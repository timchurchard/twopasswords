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
