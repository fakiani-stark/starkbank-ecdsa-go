package main

import (
	"starkbank/ecdsa-go/ellipticcurve"
	// "starkbank/ecdsa-go/ellipticcurve/utils"
	"fmt"
	"math/big"
	// "reflect"
)

func main() {
	// fmt.Println(utils.OidFromHex("2a8648ce3d0201"))
	// utils.OidFromHex("2b8104000a")
	// fmt.Println(utils.OidToHex([]int{1, 2, 840, 10045, 2, 1}))
// 	fmt.Println(utils.CreatePem(`MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAELN8kW/71zcGyXLyyCpLkOA3Aslk5J+wk
// 262ccgnGhAfg7KRidSzl4N72Xj1s90EmvIzANH40y/uQzgNPW7Y/YQ==`,`-----BEGIN PUBLIC KEY-----
// {content}
// -----END PUBLIC KEY-----`))
	// fmt.Println(utils.GetTagData("12"))
	// fmt.Println(utils.Parse("30740201010420e0ef66e4b7e3015bb745f9dfe3f91274a3ead3237d52b4d5d0f57600c7d36473a00706052b8104000aa14403420004fe7bc7bbf49041aa2cf535258d0b14b8a235e59af949a6cfadbc2d6df4ffe34f2b915573abcbc08ac6c1466a5e4511178f96a883f32af11add632ef95bdabf22"))
	// teste := big.NewInt(123)
	// fmt.Println(reflect.TypeOf(teste))
	// bigint,_ := new(big.Int).SetString("61357312058762319586231985623195", 10)
	// privateKey := ellipticcurve.NewPrivateKey(ellipticcurve.Secp256k1(), bigint)
	ellipticcurve.FromDer()
}
