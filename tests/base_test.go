package tests

import (
	"testing"
	"../ellipticcurve/utils"
	"fmt"
)

// func TestBase64(t *testing.T) {
// 	str := "tony stark"

// 	encoded := utils.Base64{}.Encode(str)
// 	decoded := utils.Base64{}.Decode(encoded)

// 	if decoded != str {
// 		t.Error("Encoding/Decoding gone wrong")
// 	}
// }

func main() {
    // fmt.Print(utils.IntFromHex("a"))
	s := "a"
	i := new(big.Int)
	i.SetString(s, 16)
	fmt.Println(i) // 10
}