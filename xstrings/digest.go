package xstrings

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func SHA1(v string) string {
	d := []byte(v)
	m := sha1.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func SHA256(v string) string {
	d := []byte(v)
	m := sha256.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func SHA512(v string) string {
	d := []byte(v)
	m := sha512.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func SHA512384(v string) string {
	d := []byte(v)
	m := sha512.New384()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
