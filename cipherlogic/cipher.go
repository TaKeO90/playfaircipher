package cipherlogic

import (
	"strings"
)

const (
	//MATRIXROWLEN the length of the matrix and the it's rows and columns
	MATRIXROWLEN int = 5
	//LETTERS the alphabet letters that we use for encrypting
	LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//PLACEHOLDER LETTER THAT IS USED TO BE THE PLACEHOLDER OF REPETITIVE LETTERS.
	PLACEHOLDER = "X"
)

var (
	endSpecialWords []string
	isEncrypt       bool
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
//checkKwLetters takes all the keyword letter and put them into a map[string]bool and give for each letter true
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

//checkEmptySpace iterate each row in the matrix and check the length of each one if it's 5 or No.
func checkEmptySpace(matrix [][]string) (arr []string, index int) {
	for i, j := range matrix {
		if len(j) != MATRIXROWLEN {
			arr, index = j, i
		}
	}
	return
}

// fillTheMatrix fill in the matrix with letters
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
	//HERE WE NEED TO CHECK IF THE LETTER WE PUT INTO THE MATRIX IS ALREADY THERE OR NOT.
	passedL := make(map[string]bool)
	for i, n := range keyword {
		if !passedL[string(n)] {
			tmp = append(tmp, string(n))
		}
		passedL[string(n)] = true
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
			ns := strings.Replace(string((*str)[1]), string((*str)[1]), PLACEHOLDER, 1)
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
			str[0] = PLACEHOLDER
			str = append(str, tmp)
			wList = append(wList, str)
		}
		if i == len(word)-1 && wList[len(wList)-1][0] == fL && wList[len(wList)-1][1] == PLACEHOLDER {
			endSpecialWords = append(endSpecialWords, PLACEHOLDER, fL)
			wList = append(wList, []string{PLACEHOLDER, fL})
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

func analyseAndEncrypt(i, j, x, y int, mRows matrixRows, matrix [][]string, result *string) {
	if i == j && x != y {
		if isEncrypt {
			mRows.shiftToRight(i, x, y, result)
		} else {
			mRows.shiftToLeft(i, x, y, result)
		}
	} else if i != j && x == y {
		if isEncrypt {
			mRows.shiftToBottom(i, j, x, result)
		} else {
			mRows.shiftToUp(i, j, x, result)
		}
	} else if i != j && x != y {
		mRows.getIntersection(matrix, i, j, x, y, result)
	}
}

// check if the word or the keyword contains space in the middle
// then edit the word or the keyword and remove that space
func checkWordKw(w string) (ok bool) {
	if strings.Contains(strings.TrimSpace(w), " ") {
		ok = true
	} else {
		ok = false
	}
	return
}

func editWordKw(w *string) {
	var tmp string
	sw := strings.Fields(*w)
	for _, n := range sw {
		tmp += n
	}
	*w = tmp
}

///////

//NewMtx get the element we need for PfMatrix type
func NewMtx(keyword string, word string, encr, decr bool) (pf *PfMatrix) {
	mtx := [][]string{}
	// HERE WE NEED TO CHECK IF THE WORD OR THE KEYWORD HAVE SPACES IN THE MIDDLE OF THEM
	if checkWordKw(word) {
		editWordKw(&word)
	}
	if checkWordKw(keyword) {
		editWordKw(&keyword)
	}
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
	} else if ind1 < len(m[rowIndex].row[rowIndex])-1 {
		ind1++
	}
	if ind2 == len(m[rowIndex].row[rowIndex])-1 {
		ind2 = 0
	} else if ind2 < len(m[rowIndex].row[rowIndex])-1 {
		ind2++
	}
	*result += m[rowIndex].row[rowIndex][ind1]
	*result += m[rowIndex].row[rowIndex][ind2]
}

func (m matrixRows) shiftToLeft(rowIndex, ind1, ind2 int, result *string) {
	if ind1 == 0 {
		ind1 = MATRIXROWLEN - 1
	} else {
		ind1--
	}
	if ind2 == 0 {
		ind2 = MATRIXROWLEN - 1
	} else {
		ind2--
	}
	*result += m[rowIndex].row[rowIndex][ind1]
	*result += m[rowIndex].row[rowIndex][ind2]
}

func (m matrixRows) shiftToBottom(fstRowIndex, sndRowIndex, index int, result *string) {
	if fstRowIndex == len(m)-1 {
		fstRowIndex = 0
	} else if fstRowIndex < len(m)-1 {
		fstRowIndex++
	}
	if sndRowIndex == len(m)-1 {
		sndRowIndex = 0
	} else if sndRowIndex < len(m)-1 {
		sndRowIndex++
	}
	*result += m[fstRowIndex].row[fstRowIndex][index]
	*result += m[sndRowIndex].row[sndRowIndex][index]
}

func (m matrixRows) shiftToUp(fstRowIndex, sndRowIndex, index int, result *string) {
	if fstRowIndex == 0 {
		fstRowIndex = MATRIXROWLEN - 1
	} else {
		fstRowIndex--
	}
	if sndRowIndex == 0 {
		sndRowIndex = MATRIXROWLEN - 1
	} else {
		sndRowIndex--
	}
	*result += m[fstRowIndex].row[fstRowIndex][index]
	*result += m[sndRowIndex].row[sndRowIndex][index]
}

func (m matrixRows) getIntersection(matrix [][]string, fstRowIndex, sndRowIndex, indx1, indx2 int, result *string) {
	if len(endSpecialWords) == 0 {
		*result += m[fstRowIndex].row[fstRowIndex][indx2]
		*result += m[sndRowIndex].row[sndRowIndex][indx1]
	} else if len(endSpecialWords) == 2 {
		specRI1, specI1 := getletterIndex(endSpecialWords[0], matrix)
		specRI2, specI2 := getletterIndex(endSpecialWords[1], matrix)
		if specRI1 == fstRowIndex && specI1 == indx1 && specRI2 == sndRowIndex && specI2 == indx2 {
			*result += m[sndRowIndex].row[sndRowIndex][indx1]
			*result += m[fstRowIndex].row[fstRowIndex][indx2]
			endSpecialWords = endSpecialWords[:0]
		} else {
			*result += m[fstRowIndex].row[fstRowIndex][indx2]
			*result += m[sndRowIndex].row[sndRowIndex][indx1]
		}
	}
}

//GenMatrix generate the matrix base on the user keyword
func (p *PfMatrix) GenMatrix() {
	var check bool
	fillTheKeyword(strings.ToUpper(p.Keyword), &p.Matrix)
	arr, index := checkEmptySpace(p.Matrix)
	if len(arr) != MATRIXROWLEN {
		check = true
	} else {
		check = false
	}
	fillInTheBlank(index, check, strings.ToUpper(p.Keyword), &p.Matrix)
}

func decfinalCheck(result *string) {
	for i, n := range *result {
		if string(n) == PLACEHOLDER {
			if i != len(*result)-1 {
				if (*result)[i-1] == (*result)[i+1] {
					*result = (*result)[:i] + (*result)[i+1:]
				}
			} else if i == len(*result)-1 {
				if string(n) == PLACEHOLDER {
					*result = (*result)[:i]
				}
			}
		}
	}
}

//EncOrDec encrypt or decrpyt user's word
func (p *PfMatrix) EncOrDec() (result string) {
	//Check if p.encrypt or p.decrypt
	if p.encrypt {
		isEncrypt = true
		// encrypt the word
		mRows := newRows(p.Matrix)
		endword := wordToPairs(p.enOrdec)
		for _, n := range endword {
			fstRI, sndRI, index1, index2 := findAndReport(n, p.Matrix)
			analyseAndEncrypt(fstRI, sndRI, index1, index2, mRows, p.Matrix, &result)
		}
	} else if p.decrypt {
		isEncrypt = false
		// decrypt the word
		genMtx := newRows(p.Matrix)
		encword := wordToPairs(p.enOrdec)
		for _, n := range encword {
			fstRI, sndRI, index1, index2 := findAndReport(n, p.Matrix)
			analyseAndEncrypt(fstRI, sndRI, index1, index2, genMtx, p.Matrix, &result)
			decfinalCheck(&result)
		}
	}
	return
}
