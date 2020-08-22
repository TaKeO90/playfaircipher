package main

import (
	"./cipherlogic"
	"fmt"
)

//TODO: set up flags for the program

func main() {
	keyword := "monarchy"
	pf := cipherlogic.NewMtx(keyword, "hello", true, false)
	matrix := pf.GenMatrix()
	fmt.Println(matrix)
	wl := pf.EncOrDec()
	fmt.Println(wl)
}
