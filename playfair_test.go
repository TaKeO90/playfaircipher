package main

import (
	"testing"

	"github.com/TaKeO90/playfaircipher/cipherlogic"
)

//TODO: make another test for decrpyting

func TestPlayfaircipher(t *testing.T) {
	ciphersAndRes := []struct {
		keyword       string
		word          string
		encryptedWord string
	}{
		{"monarchy", "balloon", "IBSUPMNA"},
		{"something", "otherthing", "MSGSUEINGA"},
		{"hello world", "fairless", "IRFAOLQYQY"},
		{"play fair", "golang", "OVAYOE"},
	}

	for _, cipherAns := range ciphersAndRes {
		pf := cipherlogic.NewMtx(cipherAns.keyword, cipherAns.word, true, false)
		pf.GenMatrix()
		result := pf.EncOrDec()
		if result != cipherAns.encryptedWord {
			t.Errorf("Encrpyted word doesn't equal what we expected %s != %s\n", result, cipherAns.encryptedWord)
		} else {
			t.Logf("SUCCESS %s\n", result)
		}
	}
}
