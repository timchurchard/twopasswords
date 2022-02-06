package cmd

import (
	"bytes"
	"testing"
)

func TestAddressMain(t *testing.T) {
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
			if got := AddressMain(out); got != tt.want {
				t.Errorf("AddressMain() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("AddressMain() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
