package num

// Lcm 最小公倍数を求める
func Lcm(x, y int) int {
	// 最大公約数 * 最小公倍数 = x * y
	// Gcd(x, y) * Lcm(x, y) = x * y

	if x == 0 || y == 0 {
		return 0
	}

	return x * y / Gcd(x, y)
}
