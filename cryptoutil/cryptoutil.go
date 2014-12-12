package cryptoutil

import (
	"errors"
	secp256k1 "github.com/toxeus/go-secp256k1"
	"math"
	"math/rand"
	"time"
)

func randInt(min int, max int) uint8 {
	//THIS IS *NOT* "cryptographically random" AND IS *NOT* SECURE.
	// PLEASE USE BETTER SOURCE OF RANDOMNESS IN PRODUCTION SYSTEMS
	// FOR DEMONSTRATION PURPOSES ONLY
	rand.Seed(time.Now().UTC().UnixNano())
	return uint8(min + rand.Intn(max-min))
}

func NewPrivateKey() []byte {
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		//THIS IS *NOT* "cryptographically random" AND IS *NOT* SECURE.
		// PLEASE USE BETTER SOURCE OF RANDOMNESS IN PRODUCTION SYSTEMS
		// FOR DEMONSTRATION PURPOSES ONLY
		bytes[i] = byte(randInt(0, math.MaxUint8))
	}
	return bytes
}

func NewPublicKey(privateKey []byte) ([]byte, error) {
	var privateKey32 [32]byte
	for i := 0; i < 32; i++ {
		privateKey32[i] = privateKey[i]
	}
	secp256k1.Start()
	publicKey, success := secp256k1.Pubkey_create(privateKey32, false)
	if !success {
		return nil, errors.New("Failed to create public key from provided private key.")
	}
	secp256k1.Stop()
	return publicKey, nil
}
