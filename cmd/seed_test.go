package cmd

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/timchurchard/twopasswords/pkg"
)

func TestSeedMain(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantOut string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
		// TODO: Add test cases.
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
