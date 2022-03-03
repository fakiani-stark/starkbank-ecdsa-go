package main

import (
	"fmt"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve"
)

func main() {
	// Create private and public key (With constant secret for testing purposes)
	privateKey := ellipticcurve.NewPrivateKey(ellipticcurve.Secp256k1(), big.NewInt(2))
	publicKey := privateKey.PublicKey()

	fmt.Printf("\n\nPrivate Key PEM: " + privateKey.ToPem())
	fmt.Printf("\n\nPublic Key PEM: " + publicKey.ToPem())

	// Sign test message using private key
	signer := ellipticcurve.Ecdsa{}.Sign("Hello World!", &privateKey)
	fmt.Printf("\n\nHello World Signature: \n" + signer.ToBase64())
	verifer := ellipticcurve.Ecdsa{}.Verify("Hello World!", signer, &publicKey)
	fmt.Printf("\n\nHello World Signature Verification: \n%t", verifer)
}
