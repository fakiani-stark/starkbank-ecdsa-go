package ecdsa

import (
	"crypto/sha256"
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/math"
	"starkbank/ecdsa-go/ellipticcurve/point"
	"starkbank/ecdsa-go/ellipticcurve/privatekey"
	"starkbank/ecdsa-go/ellipticcurve/publickey"
	"starkbank/ecdsa-go/ellipticcurve/signature"
	"starkbank/ecdsa-go/ellipticcurve/utils"
)

func Sign(message string, privateKey *privatekey.PrivateKey) signature.Signature {
	hashfunc := sha256.New()
	hashfunc.Write([]byte(message))
	byteMessage := hashfunc.Sum(nil)
	numberMessage := utils.NumberFromByteString(byteMessage)
	curve := privateKey.Curve

	var randSignPoint point.Point
	r, s := big.NewInt(0), big.NewInt(0)
	for r.Cmp(big.NewInt(0)) == 0 && s.Cmp(big.NewInt(0)) == 0 {
		randNum := utils.Random{}.Between(big.NewInt(1), big.NewInt(0).Sub(curve.N, big.NewInt(1)))
		randSignPoint = math.Multiply(curve.G, randNum, curve.N, curve.A, curve.P)
		r = big.NewInt(0).Mod(randSignPoint.X, curve.N)
		slice := big.NewInt(0).Add(numberMessage, big.NewInt(0).Mul(r, privateKey.Secret))
		s = big.NewInt(0).Mod(big.NewInt(0).Mul(slice, (math.Inv(randNum, curve.N))), curve.N)
	}

	return signature.NewSignature(*r, *s)
}

func Verify(message string, Signature signature.Signature, publicKey *publickey.PublicKey) bool {
	hashfunc := sha256.New()
	hashfunc.Write([]byte(message))
	byteMessage := hashfunc.Sum(nil)
	numberMessage := utils.NumberFromByteString(byteMessage)
	curve := publicKey.Curve
	r := &Signature.R
	s := &Signature.S

	if r.Cmp(big.NewInt(1)) < 0 && r.Cmp(big.NewInt(0).Sub(curve.N, big.NewInt(1))) > 0 {
		return false
	}
	if s.Cmp(big.NewInt(1)) < 0 && s.Cmp(big.NewInt(0).Sub(curve.N, big.NewInt(1))) > 0 {
		return false
	}

	inv := math.Inv(s, curve.N)
	u1 := math.Multiply(curve.G, big.NewInt(0).Mul(numberMessage, inv), curve.N, curve.A, curve.P)
	u2 := math.Multiply(publicKey.Point, big.NewInt(0).Mod(big.NewInt(0).Mul(r, inv), curve.N), curve.N, curve.A, curve.P)
	v := math.Add(u1, u2, curve.A, curve.P)
	if v.IsAtInfinity() {
		return false
	}
	return big.NewInt(0).Mod(v.X, curve.N).Cmp(r) == 0
}
