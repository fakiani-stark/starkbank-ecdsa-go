package ellipticcurve

import (
	"crypto/sha256"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
)

type Ecdsa struct{}

func (obj Ecdsa) Sign(message string, privateKey *PrivateKey) Signature {
	hashfunc := sha256.New()
	hashfunc.Write([]byte(message))
	byteMessage := hashfunc.Sum(nil)
	numberMessage := utils.NumberFromByteString(byteMessage)
	curve := privateKey.Curve

	var randSignPoint Point
	r, s := big.NewInt(0), big.NewInt(0)
	for r.Cmp(big.NewInt(0)) == 0 && s.Cmp(big.NewInt(0)) == 0 {
		randNum := utils.Random{}.Between(big.NewInt(1), big.NewInt(0).Sub(curve.N, big.NewInt(1)))
		randSignPoint = Multiply(curve.G, randNum, curve.N, curve.A, curve.P)
		r = big.NewInt(0).Mod(randSignPoint.X, curve.N)
		slice := big.NewInt(0).Add(numberMessage, big.NewInt(0).Mul(r, privateKey.Secret))
		s = big.NewInt(0).Mod(big.NewInt(0).Mul(slice, (Inv(randNum, curve.N))), curve.N)
	}

	return NewSignature(*r, *s)
}

func (obj Ecdsa) Verify(message string, signature Signature, publicKey *PublicKey) bool {
	hashfunc := sha256.New()
	hashfunc.Write([]byte(message))
	byteMessage := hashfunc.Sum(nil)
	numberMessage := utils.NumberFromByteString(byteMessage)
	curve := publicKey.Curve
	r := &signature.r
	s := &signature.s

	if r.Cmp(big.NewInt(1)) < 0 && r.Cmp(big.NewInt(0).Sub(curve.N, big.NewInt(1))) > 0 {
		return false
	}
	if s.Cmp(big.NewInt(1)) < 0 && s.Cmp(big.NewInt(0).Sub(curve.N, big.NewInt(1))) > 0 {
		return false
	}

	inv := Inv(s, curve.N)
	u1 := Multiply(curve.G, big.NewInt(0).Mul(numberMessage, inv), curve.N, curve.A, curve.P)
	u2 := Multiply(publicKey.Point, big.NewInt(0).Mod(big.NewInt(0).Mul(r, inv), curve.N), curve.N, curve.A, curve.P)
	v := Add(u1, u2, curve.A, curve.P)
	if v.isAtInfinity() {
		return false
	}
	return big.NewInt(0).Mod(v.X, curve.N).Cmp(r) == 0
}
