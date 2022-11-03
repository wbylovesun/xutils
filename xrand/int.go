package xrand

import "math/rand"

func Int(min, max int) int {
	return rand.Intn(max-min) + min
}
