package tests

import (
	"testing"
	"math/big"
	"../ellipticcurve/utils"
)

func TestIntegerBetween(t *testing.T) {
	a := big.NewInt(1)
	b := big.NewInt(10)
	between := utils.Random{}.Between(a, b)

	if between.Cmp(a) < 0 || between.Cmp(b) > 0 {
		t.Error("Wrong interval")
	}
}
