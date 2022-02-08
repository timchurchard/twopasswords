package pkg

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestMakeAddress(t *testing.T) {
	validp2pkh, _ := hex.DecodeString("7798cd928fd23562bceb350883b5121d0b935f97d6dc36c8cc1bf5b69aa533b5")
	validp2sh, _ := hex.DecodeString("e064997c2835646a5dc63d9741ca18febaed37082a30dcca9d0fe7515317ba4a")
	validp2wpkh, _ := hex.DecodeString("1a949fb9686c3054847163c3bcd6bce90d8cb18ca81bee0dbb5c459eb396b9d4")

	type args struct {
		mnemonic string
		password string
		num      int
		script   string
	}
	tests := []struct {
		name    string
		args    args
		want    AddressResult
		wantErr bool
	}{
		{
			name: "invalid mnemonic",
			args: args{
				mnemonic: "not a valid mnemonic",
				password: "",
				num:      0,
				script:   "p2pkh",
			},
			want:    AddressResult{},
			wantErr: true,
		}, {
			name: "invalid script",
			args: args{
				mnemonic: "fall path elegant ticket remember swamp pattern reveal section hamster timber okay electric devote pudding reject square armor sick dawn drip kit thought shrug",
				password: "",
				num:      0,
				script:   "invalid",
			},
			want:    AddressResult{},
			wantErr: true,
		}, {
			name: "valid p2pkh",
			args: args{
				mnemonic: "fall path elegant ticket remember swamp pattern reveal section hamster timber okay electric devote pudding reject square armor sick dawn drip kit thought shrug",
				password: "apple8",
				num:      123456,
				script:   "p2pkh",
			},
			want: AddressResult{
				Address:        "1BQcbiEAnhfK86CSRgoJ8JP1kVqrB8kcKG",
				SecretExponent: validp2pkh,
				Wif:            "L1EC4oJS4HKaUjLW4zsTAeKCc1vY7ZWW94KX3vnvcJSXwvBmbqcm",
				Num:            123456,
				DerivationPath: "m/44'/0'/0'/0/123456",
			},
			wantErr: false,
		}, {
			name: "valid p2wpkh-p2sh",
			args: args{
				mnemonic: "fall path elegant ticket remember swamp pattern reveal section hamster timber okay electric devote pudding reject square armor sick dawn drip kit thought shrug",
				password: "apple8",
				num:      123456,
				script:   "p2wpkh-p2sh",
			},
			want: AddressResult{
				Address:        "3M1f5Ueo1VyjMSz2Agj9g1ZyQQBL27ECFS",
				SecretExponent: validp2sh,
				Wif:            "L4juErXuC7eZRGA971hCPFYSWpwacQ4AmcR8p2ShiphHyg4Mtjbf",
				Num:            123456,
				DerivationPath: "m/49'/0'/0'/0/123456",
			},
			wantErr: false,
		}, {
			name: "valid p2wpkh",
			args: args{
				mnemonic: "fall path elegant ticket remember swamp pattern reveal section hamster timber okay electric devote pudding reject square armor sick dawn drip kit thought shrug",
				password: "apple8",
				num:      123456,
				script:   "p2wpkh",
			},
			want: AddressResult{
				Address:        "bc1qu0lepml83eyscneysfkth9sml4d70vgpxwpc2w",
				SecretExponent: validp2wpkh,
				Wif:            "Kx7NzeGpfz77DFQfnFUXJKCCMN8gGNMLwzeJqwNAkPY4m4cgaSbg",
				Num:            123456,
				DerivationPath: "m/84'/0'/0'/0/123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeAddress(tt.args.mnemonic, tt.args.password, tt.args.num, tt.args.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
