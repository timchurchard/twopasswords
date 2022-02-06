package pkg

import (
	"reflect"
	"sync"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/tyler-smith/go-bip32"
)

func Test_generateFromBytes(t *testing.T) {
	type args struct {
		prvKey   *btcec.PrivateKey
		compress bool
	}
	tests := []struct {
		name             string
		args             args
		wantWif          string
		wantAddress      string
		wantSegwitBech32 string
		wantSegwitNested string
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWif, gotAddress, gotSegwitBech32, gotSegwitNested, err := generateFromBytes(tt.args.prvKey, tt.args.compress)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateFromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWif != tt.wantWif {
				t.Errorf("generateFromBytes() gotWif = %v, want %v", gotWif, tt.wantWif)
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("generateFromBytes() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
			if gotSegwitBech32 != tt.wantSegwitBech32 {
				t.Errorf("generateFromBytes() gotSegwitBech32 = %v, want %v", gotSegwitBech32, tt.wantSegwitBech32)
			}
			if gotSegwitNested != tt.wantSegwitNested {
				t.Errorf("generateFromBytes() gotSegwitNested = %v, want %v", gotSegwitNested, tt.wantSegwitNested)
			}
		})
	}
}

func TestKey_Encode(t *testing.T) {
	type fields struct {
		path     string
		bip32Key *bip32.Key
	}
	type args struct {
		compress bool
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantWif          string
		wantAddress      string
		wantSegwitBech32 string
		wantSegwitNested string
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				path:     tt.fields.path,
				bip32Key: tt.fields.bip32Key,
			}
			gotWif, gotAddress, gotSegwitBech32, gotSegwitNested, err := k.Encode(tt.args.compress)
			if (err != nil) != tt.wantErr {
				t.Errorf("Key.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWif != tt.wantWif {
				t.Errorf("Key.Encode() gotWif = %v, want %v", gotWif, tt.wantWif)
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("Key.Encode() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
			if gotSegwitBech32 != tt.wantSegwitBech32 {
				t.Errorf("Key.Encode() gotSegwitBech32 = %v, want %v", gotSegwitBech32, tt.wantSegwitBech32)
			}
			if gotSegwitNested != tt.wantSegwitNested {
				t.Errorf("Key.Encode() gotSegwitNested = %v, want %v", gotSegwitNested, tt.wantSegwitNested)
			}
		})
	}
}

func TestKey_GetPath(t *testing.T) {
	type fields struct {
		path     string
		bip32Key *bip32.Key
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				path:     tt.fields.path,
				bip32Key: tt.fields.bip32Key,
			}
			if got := k.GetPath(); got != tt.want {
				t.Errorf("Key.GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKeyManager(t *testing.T) {
	type args struct {
		bitSize    int
		passphrase string
		mnemonic   string
	}
	tests := []struct {
		name    string
		args    args
		want    *KeyManager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKeyManager(tt.args.bitSize, tt.args.passphrase, tt.args.mnemonic)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKeyManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetMnemonic(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			if got := km.GetMnemonic(); got != tt.want {
				t.Errorf("KeyManager.GetMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetPassphrase(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			if got := km.GetPassphrase(); got != tt.want {
				t.Errorf("KeyManager.GetPassphrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetSeed(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			if got := km.GetSeed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_getKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bip32.Key
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, got1 := km.getKey(tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.getKey() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("KeyManager.getKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestKeyManager_setKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		path string
		key  *bip32.Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			km.setKey(tt.args.path, tt.args.key)
		})
	}
}

func TestKeyManager_GetMasterKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    *bip32.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, err := km.GetMasterKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyManager.GetMasterKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetMasterKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetPurposeKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		purpose uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bip32.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, err := km.GetPurposeKey(tt.args.purpose)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyManager.GetPurposeKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetPurposeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetCoinTypeKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		purpose  uint32
		coinType uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bip32.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, err := km.GetCoinTypeKey(tt.args.purpose, tt.args.coinType)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyManager.GetCoinTypeKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetCoinTypeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetAccountKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		purpose  uint32
		coinType uint32
		account  uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bip32.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, err := km.GetAccountKey(tt.args.purpose, tt.args.coinType, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyManager.GetAccountKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetAccountKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetChangeKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		purpose  uint32
		coinType uint32
		account  uint32
		change   uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bip32.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, err := km.GetChangeKey(tt.args.purpose, tt.args.coinType, tt.args.account, tt.args.change)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyManager.GetChangeKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetChangeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyManager_GetKey(t *testing.T) {
	type fields struct {
		mnemonic   string
		passphrase string
		keys       map[string]*bip32.Key
		mux        sync.Mutex
	}
	type args struct {
		purpose  uint32
		coinType uint32
		account  uint32
		change   uint32
		index    uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			km := &KeyManager{
				mnemonic:   tt.fields.mnemonic,
				passphrase: tt.fields.passphrase,
				keys:       tt.fields.keys,
				mux:        tt.fields.mux,
			}
			got, err := km.GetKey(tt.args.purpose, tt.args.coinType, tt.args.account, tt.args.change, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyManager.GetKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyManager.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
