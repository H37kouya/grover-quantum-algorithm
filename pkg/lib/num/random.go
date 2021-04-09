package num

import (
	"crypto/rand"
	"math/big"
)

// RandomInt Int returns a uniform random value in [0, max). It panics if max <= 0.
func RandomInt(num int64) (*big.Int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(num))
	if err != nil {
		return nil, err
	}

	return n, nil
}
