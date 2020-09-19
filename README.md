[![Go Report Card](https://goreportcard.com/badge/github.com/TaKeO90/playfaircipher)](https://goreportcard.com/report/github.com/TaKeO90/playfaircipher)

# [playfair cipher](https://en.wikipedia.org/wiki/Playfair_cipher)

# ROAD MAP
- [X] generate matrix with the keyword and alphabet letters on it.
- [X] split the word that we need to encrypt into pairs.
- [X] be able to get the word that the user want to encrypt and encrpyt it .
- [X] setup flags for the program to get user input for cli.
- [X] Be able to encrypted word with playfair cipher.
- [X] setup unit testing for encrpytion.
- [x] Be able to decrypt word with playfair cipher.
- [X] setup unit testing for decryption.


## Quick Start

```console
$ go mod download
$ go build .
$ ./playfaircipher -word <word to encrypt> -keyword <your keyword> -encrypt
$ ./playfaircipher -word <word to decrypt> -keyword <your keyword> -decrypt
```
**To Run unit testing**
- Run the command bellow from the root of the project .
```console
$ go test
```
