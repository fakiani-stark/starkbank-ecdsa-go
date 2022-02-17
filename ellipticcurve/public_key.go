package ellipticcurve

import (
	// "math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
	"fmt"
)

type PublicKey struct {
	Point Point
	Curve CurveFp
}

func NewPublicKey(point Point, curve CurveFp) *PublicKey {
	publicKey := new(PublicKey)
	publicKey.Curve = curve
	publicKey.Point = point
	return publicKey
}

func (self PublicKey) ToString(encoded bool) string {
	baseLength := 2 * self.Curve.Length()
	stringTemplate := fmt.Sprint("%0", baseLength, "s")
	xHex := fmt.Sprintf(stringTemplate, utils.HexFromInt(self.Point.X))
	yHex := fmt.Sprintf(stringTemplate, utils.HexFromInt(self.Point.Y))
	str := fmt.Sprint(xHex, yHex)
	fmt.Println("x: ", xHex)
	fmt.Println("y: ", yHex)
	fmt.Println("str: ", str)
	if encoded {
		return fmt.Sprint("0004", str)
	}
    return str
}

func (self PublicKey) ToDer() []byte {
	hexadecimal := utils.EncodeConstructed(
		utils.EncodeConstructed(
			utils.EncodePrimitive(utils.Object, []int64{1, 2, 840, 10045, 2, 1}),
			utils.EncodePrimitive(utils.Object, self.Curve.Oid),
		),
		utils.EncodePrimitive(utils.BitString, self.ToString(true)),
	)
	return utils.ByteStringFromHex(hexadecimal)
}

func (self PublicKey) ToPem() string {
	der := self.ToDer()
	return utils.CreatePem(utils.Base64FromByteString(der), publicKeyTemplate)
}

// func FromPem(pem string) PublicKey {
// 	publicKeyPem := utils.GetPemContent(pem, publicKeyTemplate)
// 	return FromDer()
// }

// func FromDer(data []byte) string {
// 	hexadecimal := utils.HexFromByteString(data)
// 	fmt.Println(utils.Parse(hexadecimal)[0])
// 	return ""
// }

func FromString(str string, curve CurveFp) PublicKey {
	baseLength := 2 * curve.Length()
}

const publicKeyTemplate = `
-----BEGIN PUBLIC KEY-----
{content}
-----END PUBLIC KEY-----
`
