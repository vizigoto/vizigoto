package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

// New creates a new uuid version 4 based on random numbers (see RFC 4122)
func New() string {
	bytes := make([]byte, 16)
	safeRandom(bytes)
	bytes[6] = (4 << 4) | (bytes[6] & 0xf)
	bytes[8] = bytes[8] & 0x3f
	bytes[8] = bytes[8] | 0x80
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], bytes[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], bytes[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], bytes[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], bytes[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], bytes[10:])
	return string(buf)
}

func safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}
