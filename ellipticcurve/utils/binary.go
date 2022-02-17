package utils

import (
	"strconv"
	"encoding/hex"
	"math/big"
	"fmt"
	b64 "encoding/base64"
	// "reflect"
)

// func convert(val string, base, toBase int) (string, error) {
// 	base64, err := strconv.ParseInt(val, base, 64)
// 	if err != nil {
// 		return "", err
// 	}

// 	return strconv.FormatInt(base64, toBase), nil
// }

// // 
// // Return the hexadecimal representation of the binary data. Every byte of data is converted into the
// // corresponding 2-digit hex representation. The resulting string is therefore twice as long as the length of data.
// // 
// // Params: 
// // 
// // - `data`: binary string
// // 
// // Returns:
// // 
// // - Hexadecimal string
// //
// func HexFromBinary(data string) string {

// 	ans,_ := convert(data, 2, 16)
// 	return ans
// }

// // 
// // Return the binary data represented by the hexadecimal string hexstr.
// // 
// // Params: 
// // - `data`: hexadecimal string
// // 
// // Returns:
// // 
// // - Binary string
// //
// func BinaryFromHex(data string) string {
// 	ans,_ := convert(data, 16, 2)
// 	return ans
// }

// // 
// // Get a number representation of a string
// // 
// // Params:
// // 
// // - `data`: string to be converted into a number
// // 
// // Returns:
// // 
// // - Number in hexadecimal base
// //
// func NumberFromString(data string) *big.Int {

// 	ans,_ := new(big.Int).SetString(hex.EncodeToString([]byte(data)), 16)
// 	return ans
// }

// // 
// // Get a string representation of a number
// // 
// // Params: 
// // - `number`: number to be converted into a string
// // 
// // Returns:
// // 
// // - Hexadecimal string
// //
// func StringFromNumber(number *big.Int) string {
// 	return number.String()
// }

func IntFromHex(hexadecimal string) *big.Int {
	bigInt,_ := new(big.Int).SetString(hexadecimal, 16)
	return bigInt
}

func HexFromInt(bigInt *big.Int) string {
	hexadecimal := fmt.Sprintf("%x", bigInt)
	if len(hexadecimal) % 2 == 1 {
		hexadecimal = "0" + hexadecimal
	}
	return hexadecimal
}

func BitsFromHex(hexadecimal string) string {
	bits, _ := strconv.ParseUint(hexadecimal, 16, 32)
	stringTemplate := fmt.Sprint("%0", len(hexadecimal) * 4, "b")
    return fmt.Sprintf(stringTemplate, bits)
}

func ByteStringFromBase64(base64 string) []byte {
	return []byte(b64.StdEncoding.EncodeToString([]byte(base64)))
}

func Base64FromByteString(byteString []byte) string {
	data, _ := b64.StdEncoding.DecodeString(string(byteString))
    return fmt.Sprintf("%q", data)
}

func HexFromByteString(byteString []byte) string {
	return hex.EncodeToString(byteString)
}

func ByteStringFromHex(hexadecimal string) []byte {
	data, _ := hex.DecodeString(hexadecimal)
	return data
}
