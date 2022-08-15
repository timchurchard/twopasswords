package pkg

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestMakeSalt(t *testing.T) {
	validBasicSanity, _ := hex.DecodeString("deadbeefdeadbeef")
	validBasicSanityWant, _ := hex.DecodeString("e10dc2b22b41df3057262995bb7f20477ccaa23362c9b9210687c3c7b8adb3a7")
	invalidShort, _ := hex.DecodeString("deadbeef")

	type args struct {
		password   []byte
		iterations int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "valid basic sanity",
			args: args{
				password:   validBasicSanity,
				iterations: 12345,
			},
			want:    validBasicSanityWant,
			wantErr: false,
		},
		{
			name: "invalid length",
			args: args{
				password:   invalidShort,
				iterations: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid iterations",
			args: args{
				password:   validBasicSanity,
				iterations: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeSalt(tt.args.password, tt.args.iterations)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSalt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeSeed(t *testing.T) {
	validPassword, _ := hex.DecodeString("deadbeefdeadbeef")
	validSalt, _ := hex.DecodeString("deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
	validEntropyWant, _ := hex.DecodeString("83a5ea96742b9a656a0328e271a0e790f5d3f6487db60fab157874d53be54d74")
	invalidShort, _ := hex.DecodeString("deadbeef")

	type args struct {
		password     []byte
		salt         []byte
		iterations   int
		bits         int
		language     string
		expectedSecs int
	}
	tests := []struct {
		name    string
		args    args
		want    SeedResult
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				password:     validPassword,
				salt:         validSalt,
				bits:         256,
				iterations:   12345,
				language:     "english",
				expectedSecs: 0,
			},
			want: SeedResult{
				Entropy:  validEntropyWant,
				Mnemonic: "lonely consider pitch tribe rifle crawl pool govern tiny minimum delay capable frog will capable replace autumn flavor fun trust father verify cupboard moon",
			},
		},
		{
			name: "invalid password",
			args: args{
				password:     invalidShort,
				salt:         validSalt,
				bits:         256,
				iterations:   12345,
				language:     "english",
				expectedSecs: 0,
			},
			want:    SeedResult{},
			wantErr: true,
		},
		{
			name: "invalid iterations",
			args: args{
				password:     validPassword,
				salt:         validSalt,
				bits:         256,
				iterations:   0,
				language:     "english",
				expectedSecs: 0,
			},
			want:    SeedResult{},
			wantErr: true,
		},
		{
			name: "invalid salt",
			args: args{
				password:     validPassword,
				salt:         invalidShort,
				bits:         256,
				iterations:   12345,
				language:     "english",
				expectedSecs: 0,
			},
			want:    SeedResult{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeSeed(tt.args.password, tt.args.salt, tt.args.iterations, tt.args.bits, tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
