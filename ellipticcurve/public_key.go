package ellipticcurve

import (
	"fmt"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
)

type PublicKey struct {
	Point Point
	Curve CurveFp
}

func (obj PublicKey) ToString(encoded bool) string {
	baseLength := 2 * obj.Curve.Length()
	stringTemplate := fmt.Sprint("%0", baseLength, "s")
	xHex := fmt.Sprintf(stringTemplate, utils.HexFromInt(obj.Point.X))
	yHex := fmt.Sprintf(stringTemplate, utils.HexFromInt(obj.Point.Y))
	str := fmt.Sprint(xHex, yHex)
	if encoded {
		return fmt.Sprint("0004", str)
	}
	return str
}

func (obj PublicKey) ToDer() []byte {
	hexadecimal := utils.EncodeConstructed(
		utils.EncodeConstructed(
			utils.EncodePrimitive(utils.Object, _ecdsaPublicKeyOid),
			utils.EncodePrimitive(utils.Object, obj.Curve.Oid),
		),
		utils.EncodePrimitive(utils.BitString, obj.ToString(true)),
	)
	return utils.ByteStringFromHex(hexadecimal)
}

func (obj PublicKey) ToPem() string {
	der := obj.ToDer()
	return utils.CreatePem(utils.Base64FromByteString(der), publicKeyTemplate)
}

func (obj PublicKey) FromPem(pem string) PublicKey {
	publicKeyPem := utils.GetPemContent(pem, publicKeyTemplate)
	return obj.FromDer(utils.ByteStringFromBase64(publicKeyPem))
}

func (obj PublicKey) FromDer(data []byte) PublicKey {
	hexadecimal := utils.HexFromByteString(data)
	parsed := utils.Parse(hexadecimal)[0]
	curveData := parsed.([]interface{})[0]
	pointString := parsed.([]interface{})[1].(string)
	publicKeyOid := curveData.([]interface{})[0].([]int)
	curveOid := curveData.([]interface{})[1].([]int)

	if len(publicKeyOid) != len(_ecdsaPublicKeyOid) {
		panic("Invalid Public Key Oid")
	}
	curve := CurveByOid(curveOid)
	return obj.FromString(pointString, curve, true)
}

func (obj PublicKey) FromString(str string, curve CurveFp, validatePoint bool) PublicKey {
	baseLength := 2 * curve.Length()
	if len(str) > 2*baseLength && str[:4] == "0004" {
		str = str[4:]
	}

	xs := str[:baseLength]
	ys := str[baseLength:]

	pointG := *new(Point)
	pointG.X = utils.IntFromHex(xs)
	pointG.Y = utils.IntFromHex(ys)
	pointG.Z = big.NewInt(0)

	publicKey := PublicKey{
		Point: pointG,
		Curve: curve,
	}

	if !validatePoint {
		return publicKey
	}
	if pointG.isAtInfinity() {
		panic("Public Key point is at infinity")
	}
	if !curve.Contains(pointG) {
		panic("Points is not valid in the curve")
	}
	if !Multiply(pointG, curve.N, curve.N, curve.A, curve.P).isAtInfinity() {
		panic("Point is not at infinity")
	}
	return publicKey
}

var _ecdsaPublicKeyOid = []int64{1, 2, 840, 10045, 2, 1}

const publicKeyTemplate = `
-----BEGIN PUBLIC KEY-----
{content}
-----END PUBLIC KEY-----
`
