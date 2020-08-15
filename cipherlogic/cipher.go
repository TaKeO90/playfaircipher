package cipherlogic

import (
	"strings"
)

const (
	MATRIXROWLEN int    = 5
	LETTERS      string = "ABCDEFJHIJKLMNOPQRSTUVWXYZ"
)

type PfMatrix struct {
	Keyword string
	Matrix  [][]string
	enOrdec string
}

//AUXILIARY FUNCTIONS //
func checkKwLetters(keyword string) map[string]bool {
	m := make(map[string]bool)
	for _, k := range keyword {
		m[string(k)] = true
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
	} else if letterIndex == len(LETTERS)-1 {
		(*matrix)[len((*matrix))-1] = append((*matrix)[len((*matrix))-1], letter)
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

///////

func NewMtx(keyword string) (pf *PfMatrix) {
	mtx := [][]string{}
	pf = &PfMatrix{keyword, mtx, ""}
	return
}

func (p *PfMatrix) GenMatrix() [][]string {
	var check bool
	fillTheKeyword(strings.ToUpper(p.Keyword), &p.Matrix)
	arr, index := checkEmptySpace(p.Matrix)
	if len(arr) != 5 {
		check = true
	} else {
		check = false
	}
	fillInTheBlank(index, check, strings.ToUpper(p.Keyword), &p.Matrix)
	return p.Matrix
}
