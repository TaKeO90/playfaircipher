package main

import (
	"./cipherlogic"
	"fmt"
)

func main() {
	keyword := "monarchy"
	pf := cipherlogic.NewMtx(keyword)
	matrix := pf.GenMatrix()
	fmt.Println(matrix)
}
