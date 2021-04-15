package num

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandomInt Int returns a uniform random value in [0, max). 0 <= return < max, It panics if max <= 0.
func RandomInt(num int64) (*big.Int, error) {
	if num < 0 {
		return nil, fmt.Errorf("RandomIntは正の値を指定してください num=%v", num)
	}

	n, err := rand.Int(rand.Reader, big.NewInt(num))
	if err != nil {
		return nil, err
	}

	return n, nil
}
