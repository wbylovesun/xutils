package main

import (
	"fmt"
	"github.com/wbylovesun/xutils/xrand"
)

func main() {
	fmt.Println(xrand.Chars(32))
	fmt.Println(xrand.CustomizeChars(32, xrand.Numbers))
	fmt.Println(xrand.CustomizeChars(32, xrand.LowerCaseChars))
	fmt.Println(xrand.CustomizeChars(32, xrand.UpperCaseChars))
	fmt.Println(xrand.CustomizeChars(32, xrand.SpecialChars))
	fmt.Println(xrand.CustomizeChars(32, xrand.Numbers|xrand.SpecialChars))
	fmt.Println(xrand.LowerChars(20))
	fmt.Println(xrand.UpperChars(20))
	fmt.Println(xrand.LowerCharNumbers(20))
	fmt.Println(xrand.UpperCharNumbers(20))
}
