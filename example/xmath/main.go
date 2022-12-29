package main

import (
	"fmt"
	"github.com/wbylovesun/xutils/xmath"
	"math"
)

func main() {
	x := 1.499999999999
	fmt.Println(xmath.Round(x))
	fmt.Println(math.Round(x))
}
