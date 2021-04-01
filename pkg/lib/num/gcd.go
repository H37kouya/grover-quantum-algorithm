package num

// Gcd 最大公約数を求める - ユークリッドの互除法 x < m
func Gcd(x, m int) int {
	if m == 0 {
		return x
	}

	return Gcd(m, x%m)
}
