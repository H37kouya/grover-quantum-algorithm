package valueObject

import (
	"fmt"
	"math"
)

// MaxN Nの最大値
const MaxN int = 28

type N struct {
	value int
}

func NewN(n int) (N, error) {
	if n <= 0 || MaxN < n {
		return N{value: 0}, fmt.Errorf("nは有効な値ではありません。n: %v", n)
	}

	return N{value: n}, nil
}

// NewNByTenDecimalNumber 10進数から2^Nを求める
func NewNByTenDecimalNumber(decimalNumber int) (N, error) {
	if decimalNumber <= 1 || int(math.Pow(2, float64(MaxN))) < decimalNumber {
		return N{value: 0}, fmt.Errorf("nは有効な値ではありません。n: %v", decimalNumber)
	}

	for n := 1; n <= MaxN; n++ {
		if int(math.Pow(2, float64(n))) >= decimalNumber {
			return NewN(n)
		}
	}

	return NewN(MaxN)
}

func (n N) Get() int {
	return n.value
}

// ElementCount 要素数
func (n N) ElementCount() int {
	return int(math.Pow(2, float64(n.value)))
}
