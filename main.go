package main

import (
	"fmt"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve"
)

func main() {
	var curve ellipticcurve.CurveFp = ellipticcurve.Secp256k1()
	var privKey ellipticcurve.PrivateKey = *ellipticcurve.NewPrivateKey(curve, big.NewInt(2))
	var pubKey ellipticcurve.PublicKey = *privKey.PublicKey()

	fmt.Println("", privKey.ToPem())
	fmt.Println("", privKey.FromPem("-----BEGIN EC PRIVATE KEY-----MFUCAQEEAQKgBwYFK4EEAAqhRANCAATGBH+UQe19bTBFQG6VwHzYXHeOS4zvPKerrAm5XHCe5RrhaP6mPcM5o8WEGUZs6u739jJlMmbQ4SNkMalQz+Uq-----END EC PRIVATE KEY-----").ToPem())
	fmt.Println("", pubKey.ToPem())
	fmt.Println("", pubKey.FromPem("-----BEGIN PUBLIC KEY-----MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAExgR/lEHtfW0wRUBulcB82Fx3jkuM7zynq6wJuVxwnuUa4Wj+pj3DOaPFhBlGbOru9/YyZTJm0OEjZDGpUM/lKg==-----END PUBLIC KEY-----").ToPem())
}
