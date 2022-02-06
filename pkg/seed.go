package pkg

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
)

const (
	MinPasswordBytesLen = 6

	DefaultIterations = 100000

	MinSaltIterations = 8192

	MinPBKDF2Iterations = 2048
)

// MakeSalt takes password bytes and number of interations
// and hashes is sha256(password) -> sha256(result) -> sha256(result) until iteration number of times
// returns final digest bytes, number of seconds int and error
func MakeSalt(password []byte, iterations int) ([]byte, error) {
	if len(password) < MinPasswordBytesLen {
		return nil, fmt.Errorf("password too short %d < %d", len(password), MinPasswordBytesLen)
	}
	if iterations < MinSaltIterations {
		return nil, fmt.Errorf("salt iterations below minimum %d < %d", iterations, MinSaltIterations)
	}

	result := sha256.Sum256(password[:])
	for i := 1; i < iterations; i++ {
		result = sha256.Sum256(result[:])
	}

	return result[:], nil
}

type SeedResult struct {
	Entropy  []byte
	Mnemonic string
}

// MakeSeed takes password, salt and iterations and performs pbkdf2 hmac to produce seed entropy
// return SeedResult and error
func MakeSeed(password []byte, salt []byte, iterations int, language string) (SeedResult, error) {
	if len(password) < MinPasswordBytesLen {
		return SeedResult{}, fmt.Errorf("password too short %d < %d", len(password), MinPasswordBytesLen)
	}
	if iterations < MinPBKDF2Iterations {
		return SeedResult{}, fmt.Errorf("pbkdf2 iterations below minimum %d < %d", iterations, MinPBKDF2Iterations)
	}

	if len(salt) != 32 {
		return SeedResult{}, errors.New("salt must be 32 bytes")
	}

	// todo: max iterations is 32 bits, might need to implement this so can do progress
	result := pbkdf2.Key(password, salt, iterations, 32, sha256.New)

	// bip39.SetWordList(
	mnemonic, err := bip39.NewMnemonic(result)
	if err != nil {
		return SeedResult{}, err
	}

	return SeedResult{
		Entropy:  result,
		Mnemonic: mnemonic,
	}, nil
}
