package utils

import (
	"math/big"
	"fmt"
)

// type OID interface {
// 	OidFromHex(hexadecimal string) []int
// }

type Oid struct {}

func OidFromHex(hexadecimal string) []int {
	firstByte, remainingBytes := hexadecimal[:2], hexadecimal[2:]
	firstByteInt := IntFromHex(firstByte)
	oid := []int{int(firstByteInt.Uint64() / 40), int(firstByteInt.Uint64() % 40)}
	oidInt := int(0)
	for len(remainingBytes) > 0 {
		bt := remainingBytes[0:2]
		remainingBytes = remainingBytes[2:]
		byteInt := int(IntFromHex(bt).Uint64())
		if byteInt >= 128 {
			oidInt = byteInt - 128
			continue
		}
		oidInt = oidInt * 128 + byteInt
		oid = append(oid, oidInt)
		oidInt = int(0)
	}
	return oid
}

func OidToHex(oid []int) string {
	hexadecimal := HexFromInt(big.NewInt(int64(40 * oid[0] + oid[1])))
	var byteArray []int
	for _, oidInt :=  range oid[2:] {
		endDelta := int(0)
		for {
			byteInt := oidInt % 128 + endDelta
			oidInt = oidInt / 128
			endDelta = 128
			byteArray = append(byteArray, byteInt)
			if oidInt == 0 {
				break
			}
		}
		for i := len(byteArray) - 1; i >= 0; i-- {
			hexadecimal = fmt.Sprintf("%s%s", hexadecimal, HexFromInt(big.NewInt(int64(byteArray[i]))))
		}
		byteArray = []int{}
	}
	return hexadecimal
}

