package akr78

import (
	"errors"

	"github.com/akramov1ch/crypto/internal/blake2b"
	"github.com/akramov1ch/crypto/internal/hashingSha/sha3"
)

func Akr78(data string) (string, error) {
	if data == "" {
		return "", errors.New("data is empty")
	}
	sha3haSher := sha3.NewSHA3()
	sha3haSher.Write([]byte(data))
	hash := sha3haSher.Sum()
	blake2b := blake2b.NewBlake2b()
	blake2b.Absorb(hash)
	result := blake2b.Finalize()
	if len(result) > 32 {
		result = result[:32]
	}
	return result, nil
}

func VerifyAkr78(hash string, data string) (bool, error)  {
	hash2, err := Akr78(data)
	if err != nil {
		return false, err
	}
	return hash == hash2, nil
}
