package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
)

func main() {

	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	
	if err != nil {
		panic("Failed to generate key pair: " + err.Error())
	}

	// Convert keys to Base64
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey)
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKey)

	fmt.Println("Public Key (Base64):", publicKeyBase64)
	fmt.Println("Private Key (Base64):", privateKeyBase64)
}
