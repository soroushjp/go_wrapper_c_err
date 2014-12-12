package cryptoutil

import (
	"fmt"
	"testing"
)

func TestNewPublicKey(t *testing.T) {
	privateKey := NewPrivateKey()
	publicKey, err := NewPublicKey(privateKey)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(publicKey)
}
