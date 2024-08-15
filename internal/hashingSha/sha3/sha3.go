package sha3

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

type SHA3 struct {
	hash hash.Hash
}

// NewSHA3 returns a new SHA3 hasher
func NewSHA3() *SHA3 {
	return &SHA3{
		hash: sha3.New256(),
	}
}

// Write writes data to the SHA3 hasher
func (s *SHA3) Write(data []byte) {
	s.hash.Write(data)
}

// Sum returns the SHA3 hash of the input data
func (s *SHA3) Sum() []byte {
	return s.hash.Sum(nil)
}

// pad10star1 - Padding for the SHA3 function
func pad10star1(length, rate int) []byte {
	paddingLen := ((length + 1 + rate) / rate) * rate
	pad := make([]byte, paddingLen-length)
	pad[0] = 0x01
	pad[len(pad)-1] = 0x80
	return pad
}
