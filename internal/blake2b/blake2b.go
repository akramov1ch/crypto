package blake2b

import (
	"encoding/hex"
	"hash"
	"log"

	"golang.org/x/crypto/blake2b"
)

// Blake2b structure
type Blake2b struct {
	hash hash.Hash
}

// NewBlake2b returns a new Blake2b hasher
func NewBlake2b() *Blake2b {
	hash, err := blake2b.New256(nil)
	if err != nil {
		log.Fatalf("Failed to create Blake2b hasher: %v", err)
	}
	return &Blake2b{
		hash: hash,
	}
}

// Absorb absorbs input data into the Blake2b hash
func (b *Blake2b) Absorb(data []byte) {
	_, err := b.hash.Write(data)
	if err != nil {
		log.Fatalf("Failed to write data to Blake2b hasher: %v", err)
	}
}

// Finalize finalizes the Blake2b hash and returns the result
func (b *Blake2b) Finalize() string {
	return hex.EncodeToString(b.hash.Sum(nil))
}
