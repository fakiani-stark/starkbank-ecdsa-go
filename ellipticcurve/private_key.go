package ellipticcurve

import (
	"crypto/rand"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
)

type PrivateKey struct {
	Curve  CurveFp
	Secret *big.Int
}

func NewPrivateKey(curve CurveFp, secret *big.Int) *PrivateKey {
	privateKey := new(PrivateKey)
	privateKey.Curve = curve
	privateKey.Secret = secret
	randomSecret, err := rand.Int(rand.Reader, curve.N)
	if err != nil {
		return privateKey
	}
	if len(secret.Bits()) == 0 {
		privateKey.Secret = randomSecret
	}
	return privateKey
}

func (obj PrivateKey) PublicKey() *PublicKey {
	pointG := *new(Point)
	pointG.X = obj.Curve.Gx
	pointG.Y = obj.Curve.Gy
	calculatedPoint := Multiply(
		pointG,
		obj.Secret,
		obj.Curve.N,
		obj.Curve.A,
		obj.Curve.P,
	)
	return NewPublicKey(calculatedPoint, obj.Curve)
}

func (obj PrivateKey) ToString() string {
	return utils.HexFromInt(obj.Secret)
}

func (obj PrivateKey) ToDer() []byte {
	publicKeyString := obj.PublicKey().ToString(true)
	hexadecimal := utils.EncodeConstructed(
		utils.EncodePrimitive(utils.Integer, big.NewInt(1)),
		utils.EncodePrimitive(utils.OctetString, utils.HexFromInt(obj.Secret)),
		utils.EncodePrimitive(utils.OidContainer,
			utils.EncodePrimitive(utils.Object, obj.Curve.Oid)),
		utils.EncodePrimitive(utils.PublicKeyPointContainer,
			utils.EncodePrimitive(utils.BitString, publicKeyString)),
	)
	return utils.ByteStringFromHex(hexadecimal)
}

func (obj PrivateKey) ToPem() string {
	der := obj.ToDer()
	return utils.CreatePem(utils.Base64FromByteString(der), privateKeyTemplate)
}

func (obj PrivateKey) FromPem(pem string) *PrivateKey {
	privateKeyPem := utils.GetPemContent(pem, privateKeyTemplate)
	return obj.FromDer(utils.ByteStringFromBase64(privateKeyPem))
}

func (obj PrivateKey) FromDer(data []byte) *PrivateKey {
	hexadecimal := utils.HexFromByteString(data)
	parsed := utils.Parse(hexadecimal)[0]
	privateKeyFlag := parsed.([]interface{})[0].(*big.Int)
	secretHex := parsed.([]interface{})[1].(string)
	curveOid := parsed.([]interface{})[2].([]interface{})[0].([]int)
	publicKeyString := parsed.([]interface{})[3].([]interface{})[0].(string)
	curve := CurveByOid(curveOid)
	privateKey := obj.FromString(secretHex, curve)

	if privateKeyFlag.Cmp(big.NewInt(1)) != 0 {
		panic("Private keys should start with a '1' flag")
	}

	if privateKey.PublicKey().ToString(true) != publicKeyString {
		panic("Private keys should start with a '1' flag")
	}

	return obj.FromString(secretHex, curve)
}

func (obj PrivateKey) FromString(str string, curve CurveFp) *PrivateKey {
	return NewPrivateKey(curve, utils.IntFromHex(str))
}

const privateKeyTemplate = `
-----BEGIN EC PRIVATE KEY-----
{content}
-----END EC PRIVATE KEY-----
`
