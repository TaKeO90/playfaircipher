package main

import (
	"./cipherlogic"
	"fmt"
)

//TODO: i should skip letter j and give it the same index as i

func main() {
	keyword := "monarchy"
	pf := cipherlogic.NewMtx(keyword, "BALLOON", true, false)
	matrix := pf.GenMatrix()
	fmt.Println(matrix)
	wl := pf.EncOrDec()
	fmt.Println(wl)
}
