package ellipticcurve

import (
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/utils"
)

type Signature struct {
	r big.Int
	s big.Int
}

func NewSignature(R big.Int, S big.Int) Signature {
	return Signature{
		r: R,
		s: S,
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
		utils.EncodePrimitive(utils.Integer, &obj.r),
		utils.EncodePrimitive(utils.Integer, &obj.s),
	)
}

func SignatureFromBase64(str string) Signature {
	der := utils.ByteStringFromBase64(str)
	return SignatureFromDer(der)
}

func SignatureFromDer(str []byte) Signature {
	hexadecimal := utils.HexFromByteString(str)
	return SignatureFromString(hexadecimal)
}

func SignatureFromString(str string) Signature {
	parse := utils.Parse(str)[0]
	r := parse.([]interface{})[0].(*big.Int)
	s := parse.([]interface{})[1].(*big.Int)
	return NewSignature(*r, *s)
}
