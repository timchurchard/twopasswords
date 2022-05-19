package cmd

import (
	"bytes"
	"encoding/hex"
	"flag"
	"os"
	"reflect"
	"testing"

	"github.com/timchurchard/twopasswords/pkg"
)

func TestSeedMain(t *testing.T) {
	const cliName = "seed"

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
				"helloworld",
				"--iterations",
				"123456",
			},
			want:    0,
			wantOut: "Made seed. Hex = 0d597f6ab5642fe586ef7ae8683cb3c3b515510aa98f5710af5d8737a3576732\nMnemonic = ask slogan survey hello drill version brick urban trick draft coconut manual eyebrow possible click cradle finish lyrics student attack kick produce orphan divorce\n",
		},
	}
	for _, tt := range tests {
		// reset flags else panic
		flag.CommandLine = flag.NewFlagSet(cliName, flag.ExitOnError)
		os.Args = append([]string{cliName}, tt.args...)

		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if got := SeedMain(out); got != tt.want {
				t.Errorf("SeedMain() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("SeedMain() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_makeSeed(t *testing.T) {
	sanityEntropy, _ := hex.DecodeString("df43c0019f4c8a04132779cc871b2645b0107f8452770ed4505976e0269efa03")

	type args struct {
		password   string
		iterations int
	}
	tests := []struct {
		name    string
		args    args
		want    pkg.SeedResult
		wantOut string
		wantErr bool
	}{
		{
			name: "sanity",
			args: args{
				password:   "apple8",
				iterations: 12345,
			},
			want: pkg.SeedResult{
				Entropy:  sanityEntropy,
				Mnemonic: "tent bulk about direct silly acoustic erode upset smart decide sister merge absurd divert bacon exclude attract penalty bind universe act exhaust trend improve",
			},
			wantOut: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := makeSeed(out, tt.args.password, tt.args.iterations)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSeed() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("makeSeed() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
