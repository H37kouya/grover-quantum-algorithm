package lib

import (
	"crypto/rand"
	"math/big"
)

func RandomInt(num int64) (*big.Int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(num))
	if err != nil {
		return nil, err
	}

	return n, nil
}
