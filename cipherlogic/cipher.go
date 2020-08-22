package cipherlogic

import (
	"strings"
)

const (
	//MATRIXROWLEN the length of the matrix and the it's rows and columns
	MATRIXROWLEN int = 5
	//LETTERS the alphabet letters that we use for encrypting
	LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//PfMatrix holds elements that we need for the encryption or decryption
// or can help us figure out what the user need `encrypt or decrypt`
type PfMatrix struct {
	Keyword string
	Matrix  [][]string
	enOrdec string
	encrypt bool
	decrypt bool
}

type matrixR struct {
	row map[int][]string
}

type matrixRows []matrixR

//AUXILIARY FUNCTIONS //
func checkKwLetters(keyword string) map[string]bool {
	m := make(map[string]bool)
	for _, k := range keyword {
		m[string(k)] = true
	}
	if !m["J"] {
		m["J"] = true
	}
	return m
}

func checkEmptySpace(matrix [][]string) (arr []string, index int) {
	for i, j := range matrix {
		if len(j) != MATRIXROWLEN {
			arr, index = j, i
		}
	}
	return
}

func fillTheMatrix(s *[]string, matrix *[][]string, letterIndex int, letter string) {
	if len(*s) == MATRIXROWLEN {
		(*matrix) = append((*matrix), (*s))
		*s = []string{}
	}
}

func fillInTheBlank(index int, check bool, keyword string, matrix *[][]string) {
	m := checkKwLetters(keyword)
	s := []string{}
	passed := make(map[string]bool)
	for letterIndex, letter := range LETTERS {
		if !m[string(letter)] {
			if check {
				if len((*matrix)[index]) != MATRIXROWLEN {
					passed[string(letter)] = true
					(*matrix)[index] = append((*matrix)[index], string(letter))
				}
				if !passed[string(letter)] {
					s = append(s, string(letter))
				}
				fillTheMatrix(&s, matrix, letterIndex, string(letter))
			} else {
				s = append(s, string(letter))
				fillTheMatrix(&s, matrix, letterIndex, string(letter))
			}
		}
	}
}

func fillTheKeyword(keyword string, matrix *[][]string) {
	tmp := []string{}
	for i, n := range keyword {
		tmp = append(tmp, string(n))
		if len(tmp) == MATRIXROWLEN {
			(*matrix) = append((*matrix), tmp)
			tmp = []string{}
		} else if len(tmp) != MATRIXROWLEN && i == len(keyword)-1 {
			(*matrix) = append((*matrix), tmp)
		}
	}
}

func isReplicated(str *[]string, fL *string) {
	if len(*str) == 2 {
		cmp := string((*str)[0])
		if string((*str)[1]) == cmp {
			tmp := []string{}
			tmp = append(tmp, cmp)
			ns := strings.Replace(string((*str)[1]), string((*str)[1]), "X", 1)
			tmp = append(tmp, ns)
			*str = tmp
			*fL = string((*str)[0])
		}
	}
}

func wordToPairs(word string) (wList [][]string) {
	var str []string
	var fL string
	for i, w := range word {
		if fL != "" {
			str = append(str, fL, string(w))
			fL = ""
		} else {
			str = append(str, string(w))
		}
		if len(str) == 2 {
			isReplicated(&str, &fL)
			wList = append(wList, str)
			str = []string{}
		} else if i == len(word)-1 && len(str) != 2 {
			tmp := str[0]
			str[0] = "X"
			str = append(str, tmp)
			wList = append(wList, str)
		}
	}
	return
}

func isIn(w1, w2 string, list []string) (indx1, indx2 int, found1, found2 bool) {
	for i, j := range list {
		if w1 == j {
			indx1 = i
			found1 = true
		} else if w2 == j {
			indx2 = i
			found2 = true
		}
	}
	return
}

func getletterIndex(letter string, matrix [][]string) (rowIndex, index int) {
	for i, r := range matrix {
		for j, l := range r {
			if l == letter {
				rowIndex = i
				index = j
			}
		}
	}
	return
}

func findAndReport(pair []string, matrix [][]string) (fstRowIndex, sndRowIndex int, indx1, indx2 int) {
	//var res []int
	//var rowIndex []int
	for b, w := range matrix {
		index1, index2, found1, found2 := isIn(pair[0], pair[1], w)
		if found1 {
			indx1 = index1
			fstRowIndex = b
		} else if !found1 && pair[0] == "J" {
			x, y := getletterIndex("I", matrix)
			indx1 = y
			fstRowIndex = x
		}
		if found2 {
			indx2 = index2
			sndRowIndex = b
		} else if !found2 && pair[1] == "J" {
			x, y := getletterIndex("I", matrix)
			indx2 = y
			sndRowIndex = x
		}
	}
	return
}

func analyseAndEncrypt(i, j, x, y int, mRows matrixRows, result *string) {
	if i == j && x != y {
		mRows.shiftToRight(i, x, y, result)
	} else if i != j && x == y {
		mRows.shiftToBottom(i, j, x, result)
	} else if i != j && x != y {
		mRows.getIntersection(i, j, x, y, result)
	}
}

///////

//NewMtx get the element we need for PfMatrix type
func NewMtx(keyword string, word string, encr, decr bool) (pf *PfMatrix) {
	mtx := [][]string{}
	pf = &PfMatrix{keyword, mtx, strings.ToUpper(word), encr, decr}
	return
}

func newRows(matrix [][]string) matrixRows {
	var mRows matrixRows
	mtx := new(matrixR)
	for i, r := range matrix {
		m := make(map[int][]string)
		m[i] = r
		mtx.row = m
		mRows = append(mRows, *mtx)
	}
	return mRows
}

func (m matrixRows) shiftToRight(rowIndex, ind1, ind2 int, result *string) {
	if ind1 == len(m[rowIndex].row[rowIndex])-1 {
		ind1 = 0
	}
	if ind2 == len(m[rowIndex].row[rowIndex])-1 {
		ind2 = 0
	}
	if ind1 < len(m[rowIndex].row[rowIndex])-1 {
		ind1++
	}
	if ind2 < len(m[rowIndex].row[rowIndex])-1 {
		ind2++
	}
	*result += m[rowIndex].row[rowIndex][ind1]
	*result += m[rowIndex].row[rowIndex][ind2]
}

func (m matrixRows) shiftToBottom(fstRowIndex, sndRowIndex, index int, result *string) {
	if fstRowIndex == len(m)-1 {
		fstRowIndex = 0
	}
	if sndRowIndex == len(m)-1 {
		sndRowIndex = 0
	}
	if fstRowIndex < len(m)-1 {
		fstRowIndex++
	}
	if sndRowIndex < len(m)-1 {
		sndRowIndex++
	}
	*result += m[fstRowIndex].row[fstRowIndex][index]
	*result += m[sndRowIndex].row[sndRowIndex][index]
}

func (m matrixRows) getIntersection(fstRowIndex, sndRowIndex, indx1, indx2 int, result *string) {
	*result += m[fstRowIndex].row[fstRowIndex][indx2]
	*result += m[sndRowIndex].row[sndRowIndex][indx1]
}

//GenMatrix generate the matrix base on the user keyword
func (p *PfMatrix) GenMatrix() [][]string {
	var check bool
	fillTheKeyword(strings.ToUpper(p.Keyword), &p.Matrix)
	arr, index := checkEmptySpace(p.Matrix)
	if len(arr) != MATRIXROWLEN {
		check = true
	} else {
		check = false
	}
	fillInTheBlank(index, check, strings.ToUpper(p.Keyword), &p.Matrix)
	return p.Matrix
}

//EncOrDec encrypt or decrpyt user's word
func (p *PfMatrix) EncOrDec() (result string) {
	//Check if p.encrypt or p.decrypt
	if p.encrypt {
		// encrypt the word
		mRows := newRows(p.Matrix)
		endword := wordToPairs(p.enOrdec)
		for _, n := range endword {
			fstRI, sndRI, index1, index2 := findAndReport(n, p.Matrix)
			analyseAndEncrypt(fstRI, sndRI, index1, index2, mRows, &result)
		}
	} else if p.decrypt {
		// decrypt the word
		result = ""
	}
	return
}
