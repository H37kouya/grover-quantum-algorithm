package num

import "math"

// Round 四捨五入
func Round(num, places float64) float64 {
	shift := math.Pow(10, places)

	return roundInt(num*shift) / shift
}

// RoundUp 切り上げ
func RoundUp(num, places float64) float64 {
	shift := math.Pow(10, places)

	return roundUpInt(num*shift) / shift
}

// RoundDown 切り捨て
func RoundDown(num, places float64) float64 {
	shift := math.Pow(10, places)

	return math.Trunc(num*shift) / shift
}

// roundInt 四捨五入(整数)
func roundInt(num float64) float64 {
	t := math.Trunc(num)
	if math.Abs(num-t) >= 0.5 {
		return t + math.Copysign(1, num)
	}

	return t
}

// roundInt 切り上げ(整数)
func roundUpInt(num float64) float64 {
	t := math.Trunc(num)

	return t + math.Copysign(1, num)
}
