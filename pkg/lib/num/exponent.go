package num

import (
	"fmt"
	"math"
)

// Exponent 指数型
type Exponent struct {
	Base int // 底
	Exp  int // 指数
}

// CalcExponent べき乘の形をもとめる
func CalcExponent(num int) (Exponent, error) {
	size := num - 2

	// 3以下のとき
	if num <= 3 {
		return Exponent{Base: 0, Exp: 0}, fmt.Errorf("べき乗の形ではありませんでした。")
	}

	// 2の累乗のとき
	if num&(num-1) == 0 {
		return Exponent{Base: 2, Exp: int(math.Log(float64(num)) / math.Ln2)}, nil
	}

	// 3の累乗のとき
	for base := 3; true; base++ {
		exp := 2

		for exp = 2; exp < size; exp++ {
			pow := math.Pow(float64(base), float64(exp))

			if pow == float64(num) {
				return Exponent{Base: base, Exp: exp}, nil
			}

			if pow > float64(num) {
				break
			}
		}

		if exp == 2 {
			break
		}
	}

	return Exponent{Base: 0, Exp: 0}, fmt.Errorf("べき乗の形ではありませんでした。")
}
