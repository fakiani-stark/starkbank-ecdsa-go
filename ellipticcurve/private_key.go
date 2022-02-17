package ellipticcurve

import (
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
	// "fmt"
)

type PrivateKey struct {
	Curve CurveFp
	Secret *big.Int
}

func NewPrivateKey(curve CurveFp, secret *big.Int) *PrivateKey {
	privateKey := new(PrivateKey)
	privateKey.Curve = curve
	privateKey.Secret = secret
	return privateKey
}

// func PublicKey()

func (self PrivateKey) ToString() string {
	return utils.HexFromInt(self.Secret)
}

// func (self PrivateKey) ToDer() {
	
// 	ByteStringFromHex
// }

// func (self PrivateKey) ToPem() string {
// 	der := ToDer()
// }

// func FromPem(pem string) PrivateKey {
// 	privateKeyPem := utils.GetPemContent(pem, privateKeyTemplate)
// 	return FromDer()
// }

// func FromDer(data []byte) string {
// 	hexadecimal := utils.HexFromByteString(data)
// 	fmt.Println(utils.Parse(hexadecimal)[0])
// 	return ""
// }

func FromString(str string, curve CurveFp) PrivateKey {
	return *NewPrivateKey(curve, utils.IntFromHex(str))
}

const privateKeyTemplate = `
-----BEGIN EC PRIVATE KEY-----
{content}
-----END EC PRIVATE KEY-----
`
