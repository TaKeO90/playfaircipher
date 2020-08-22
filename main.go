package main

import (
	"flag"
	"fmt"

	"github.com/TaKeO90/playfaircipher/cipherlogic"
)

func flagParser(keyword, word *string, encrypt, decrypt *bool) {
	flag.StringVar(keyword, "keyword", "", "specify the keyword that will generate a new matrix each time you change this keyword")
	flag.StringVar(word, "word", "", "the word you want to encrpyt or decrypt")
	flag.BoolVar(encrypt, "encrypt", false, "use this flag if you want to encrypt the word")
	flag.BoolVar(decrypt, "decrypt", false, "use this flag if you want to decrypt the word")
	flag.Parse()
}

func main() {
	var (
		encrypt bool
		decrypt bool
		keyword string
		word    string
	)
	flagParser(&keyword, &word, &encrypt, &decrypt)

	if keyword != "" && word != "" && encrypt && !decrypt {
		pf := cipherlogic.NewMtx(keyword, word, encrypt, decrypt)
		pf.GenMatrix()
		result := pf.EncOrDec()
		fmt.Println(result)
	} else if keyword != "" && word != "" && decrypt && !encrypt {
		pf := cipherlogic.NewMtx(keyword, word, encrypt, decrypt)
		pf.GenMatrix()
		result := pf.EncOrDec()
		fmt.Println(result)
	} else {
		flag.PrintDefaults()
	}
}
