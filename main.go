package main

import (
	"fmt"

	"github.com/viktor-br/ds/problems"

	"math/big"
)

func main() {
	problems.IsUniqueUtf8("фбфкд")

	a := big.NewInt(0)

	a.SetBit(a, 70, 1)

	fmt.Println(a.String())

	// fmt.Printf("%d %b %b\n", a|b, a, b)
}
