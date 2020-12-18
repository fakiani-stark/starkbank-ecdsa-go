package utils

import (
	"strconv"
	"encoding/hex"
	"math/big"
)

func convert(val string, base, toBase int) (string, error) {
	base64, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(base64, toBase), nil
}

func HexFromBinary(data string) string {
	// 
	// Return the hexadecimal representation of the binary data. Every byte of data is converted into the
	// corresponding 2-digit hex representation. The resulting string is therefore twice as long as the length of data.
	// 
	// Args: 
	// data: binary string
	// 
	// Returns:
	// Hexadecimal string
	//
	ans,_ := convert(data, 2, 16)
	return ans
}

func BinaryFromHex(data string) string {
	// 
	// Return the binary data represented by the hexadecimal string hexstr.
	// 
	// Args: 
	// data: hexadecimal string
	// 
	// Returns:
	// Binary string
	//
	ans,_ := convert(data, 16, 2)
	return ans
}

func NumberFromString(data string) *big.Int {
	// 
	// Get a number representation of a string
	// 
	// Args:
	// data: string to be converted into a number
	// 
	// Returns:
	// Number in hexadecimal base
	//
	ans,_ := new(big.Int).SetString(hex.EncodeToString([]byte(data)), 16)
	return ans
}

func StringFromNumber(number *big.Int) string {
	// 
	// Get a string representation of a number
	// 
	// Args: 
	// number: number to be converted into a string
	// 
	// Returns:
	// Hexadecimal string
	//
	return number.String()
}