package main

import (
	"fmt"
	"github.com/soroushjp/go_wrapper_c_err/cryptoutil"
	"log"
)

func main() {
	privateKey := cryptoutil.NewPrivateKey()
	publicKey, err := cryptoutil.NewPublicKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(publicKey)
}
