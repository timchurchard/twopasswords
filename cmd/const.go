package cmd

const (
	defaultEmpty  = ""
	default256    = 256
	defaultScript = "p2wpkh"

	usageIterations = "Number of iterations for PBKDF2"
	usageBits       = "Number of bits for seed. Must be 128 or 256 (default)"
	usagePassword   = "Password for seed"

	usageSecond = "Password for bip39 and wallet encryption"
	usageNum    = "Address number to make"
	usageScript = "Script (p2pkh, p2wpkh, p2wpkh-p2sh)"
	usageBip38  = "Encrypt private key with bip38 passphrase"
)
