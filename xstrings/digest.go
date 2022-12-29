package xstrings

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
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

func CRC32(v string) uint32 {
	return crc32.ChecksumIEEE([]byte(v))
}

func SHA256WithHMAC(v, k string) ([]byte, error) {
	hash := hmac.New(sha256.New, []byte(k))
	_, err := hash.Write([]byte(v))
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func SHA256WithHMAC2Hex(v, k string) (string, error) {
	b, e := SHA256WithHMAC(v, k)
	if e != nil {
		return "", e
	}
	return hex.EncodeToString(b), nil
}

func SHA256WithHMAC2Base64Str(v, k string) (string, error) {
	b, e := SHA256WithHMAC(v, k)
	if e != nil {
		return "", e
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
