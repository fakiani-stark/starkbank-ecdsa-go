package signature

import (
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
)

type Signature struct {
	R big.Int
	S big.Int
}

func NewSignature(r big.Int, s big.Int) Signature {
	return Signature{
		R: r,
		S: s,
	}
}

func (obj Signature) ToDer() []byte {
	hexadecimal := obj.ToString()
	return utils.ByteStringFromHex(hexadecimal)
}

func (obj Signature) ToBase64() string {
	return utils.Base64FromByteString(obj.ToDer())
}

func (obj Signature) ToString() string {
	return utils.EncodeConstructed(
		utils.EncodePrimitive(utils.Integer, &obj.R),
		utils.EncodePrimitive(utils.Integer, &obj.S),
	)
}

func FromBase64(str string) Signature {
	der := utils.ByteStringFromBase64(str)
	return FromDer(der)
}

func FromDer(str []byte) Signature {
	hexadecimal := utils.HexFromByteString(str)
	return FromString(hexadecimal)
}

func FromString(str string) Signature {
	parse := utils.Parse(str)[0]
	r := parse.([]interface{})[0].(*big.Int)
	s := parse.([]interface{})[1].(*big.Int)
	return NewSignature(*r, *s)
}
