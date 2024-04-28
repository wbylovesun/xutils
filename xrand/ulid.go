package xrand

import (
	"github.com/oklog/ulid/v2"
)

func ULID() string {
	// uLid生成的默认是26位的
	newUlid := ulid.Make()
	return newUlid.String()
}
