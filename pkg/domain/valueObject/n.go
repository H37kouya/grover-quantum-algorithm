package valueObject

import (
	"fmt"
	"math"
)

type N struct {
	value int
}

func NewN(n int) (N, error) {
	if n <= 0 || 28 < n {
		return N{value: 0}, fmt.Errorf("nは有効な値ではありません。n: %v", n)
	}

	return N{value: n}, nil
}

func (n N) Get() int {
	return n.value
}

// ElementCount 要素数
func (n N) ElementCount() int {
	return int(math.Pow(2, float64(n.value)))
}
