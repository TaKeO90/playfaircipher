package main

import (
	"./cipherlogic"
	"fmt"
)

func main() {
	keyword := "monarchy"
	pf := cipherlogic.NewMtx(keyword, "BALLOON", true, false)
	matrix := pf.GenMatrix()
	fmt.Println(matrix)
	wl := pf.EncOrDec()
	fmt.Println(wl)
}
