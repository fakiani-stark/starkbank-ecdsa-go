package main

import (
	"fmt"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve"
)

func main() {
	// Create private and public key (With constant secret for testing purposes)
	privateKey := ellipticcurve.NewPrivateKey(ellipticcurve.Secp256k1(), big.NewInt(0))
	publicKey := privateKey.PublicKey()

	fmt.Printf("\n\nPrivate Key PEM: " + privateKey.ToPem())
	fmt.Printf("\n\nPublic Key PEM: " + publicKey.ToPem())

	// Signing and Verification
	testString := "Hello World!"
	signer := ellipticcurve.Ecdsa{}.Sign(testString, &privateKey)
	verifer := ellipticcurve.Ecdsa{}.Verify(testString, signer, &publicKey)
	fmt.Printf("\n\n'%s' Signature & Verification: \n%s\n%t\n\n", testString, signer.ToBase64(), verifer)
}
