package xrand

import (
	"math/rand"
	"time"
)

const (
	Numbers        = 1
	LowerCaseChars = 2
	UpperCaseChars = 4
	SpecialChars   = 8
)

var numbers = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}
var upperChars = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G',
	'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}
var lowerChars = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}
var specialChars = []byte{
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'-', '_', '+', '=', '/', '?', '<', '>', ',', '.',
	'[', ']', '|', '~', ';', ':', '"', '\'', '{', '}',
	'`',
}

var upperCharNumbers = append(numbers, upperChars...)
var lowerCharNumbers = append(numbers, lowerChars...)
var allChars = append(upperCharNumbers, lowerChars...)

var charMap = map[int][]byte{
	Numbers:                                                  numbers,
	LowerCaseChars:                                           lowerChars,
	UpperCaseChars:                                           upperChars,
	SpecialChars:                                             specialChars,
	Numbers | LowerCaseChars:                                 lowerCharNumbers,
	Numbers | UpperCaseChars:                                 upperCharNumbers,
	Numbers | SpecialChars:                                   append(numbers, specialChars...),
	LowerCaseChars | UpperCaseChars:                          append(lowerChars, upperChars...),
	LowerCaseChars | SpecialChars:                            append(lowerChars, specialChars...),
	UpperCaseChars | SpecialChars:                            append(upperChars, specialChars...),
	Numbers | LowerCaseChars | UpperCaseChars:                append(lowerChars, upperCharNumbers...),
	Numbers | LowerCaseChars | SpecialChars:                  append(lowerCharNumbers, specialChars...),
	Numbers | UpperCaseChars | SpecialChars:                  append(upperCharNumbers, specialChars...),
	LowerCaseChars | UpperCaseChars | SpecialChars:           append(append(lowerChars, upperChars...), specialChars...),
	Numbers | LowerCaseChars | UpperCaseChars | SpecialChars: append(allChars, specialChars...),
}

// CustomizeChars 以指定的集合生成指定长度的随机串
//
// cs: Numbers | LowerCaseChars | UpperCaseChars | SpecialChars 集合
func CustomizeChars(l int, cs int) string {
	if l <= 0 {
		return ""
	}
	chrs, ok := charMap[cs]
	if !ok {
		return ""
	}
	return charNumbers(chrs, l)
}

func Chars(l int) string {
	return charNumbers(allChars, l)
}

func LowerChars(l int) string {
	return charNumbers(lowerChars, l)
}

func UpperChars(l int) string {
	return charNumbers(upperChars, l)
}

func LowerCharNumbers(l int) string {
	return charNumbers(lowerCharNumbers, l)
}

func UpperCharNumbers(l int) string {
	return charNumbers(upperCharNumbers, l)
}

func charNumbers(charCollection []byte, l int) string {
	if l <= 0 {
		return ""
	}
	cl := len(charCollection)
	var r []byte
	for i := 0; i < l; i += 1 {
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(cl)
		r = append(r, charCollection[k])
	}
	return string(r)
}
