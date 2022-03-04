package main

import (
	"fmt"
	"starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"starkbank/ecdsa-go/ellipticcurve/privatekey"
)

func main() {
	// Generate privateKey from PEM string
	privateKey := privatekey.FromPem(
		"-----BEGIN EC PARAMETERS-----" +
			"BgUrgQQACg==" +
			"-----END EC PARAMETERS-----" +
			"-----BEGIN EC PRIVATE KEY-----" +
			"MHQCAQEEIODvZuS34wFbt0X53+P5EnSj6tMjfVK01dD1dgDH02RzoAcGBSuBBAAK" +
			"oUQDQgAE/nvHu/SQQaos9TUljQsUuKI15Zr5SabPrbwtbfT/408rkVVzq8vAisbB" +
			"RmpeRREXj5aog/Mq8RrdYy75W9q/Ig==" +
			"-----END EC PRIVATE KEY-----")

	message := `
    "transfers": [
        {
            "amount": 100000000,
            "taxId": "594.739.480-42",
            "name": "Daenerys Targaryen Stormborn",
            "bankCode": "341",
            "branchCode": "2201",
            "accountNumber": "76543-8",
            "tags": ["daenerys", "targaryen", "transfer-1-external-id"]
        }
    ]`

	signature := ecdsa.Sign(message, &privateKey)

	// Generate Signature in base64. This result can be sent to Stark Bank in the request header as the Digital-Signature parameter.
	fmt.Println(signature.ToBase64())

	// To double check if the message matches the signature, do this:
	publicKey := privateKey.PublicKey()

	fmt.Println(ecdsa.Verify(message, signature, &publicKey))
}
