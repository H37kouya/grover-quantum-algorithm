package service

import (
	"fmt"
	"grover-quantum-search/pkg/lib/num"
	"math"
	"sort"
)

// ShorFactorization ショアのアルゴリズム
func ShorFactorization(M int) ([]int, error) {
	fmt.Println("ShorFactorization args", M)

	// (1) もしMが偶数なら、素因数2を出力する
	// (2) M = a^b (a >= 1, b >= 2)かどうかを確かめる
	// もしそうであれば、素因数aを出力する
	// (3) 1 から M - 1 の間で任意に x を選ぶ。
	// もし x と M の最大公約数 (gcd(x, M)) が 1 より大きい (gcd(x, M) > 1)ならば、gcd(x, M)を出力する
	// (4) x, M (x < M) の位数 r を計算する (x^r mod M = 1)
	// (5) もし r が偶数であり、x^(r/2) mod M != 1 ならば、gcd(x^(r/2) - 1, M)とgcd(x^(r/2) + 1, M)を計算する
	// もしこれらのうち一つが M の因数なら、それを出力する。だめなら(3)へ戻る。

	// 自然数であるかを事前に確認する
	if M <= 0 {
		return []int{}, fmt.Errorf("引数Mは自然数でなかればなりません。M=%v", M)
	}

	// 1は事前にはじく
	if M == 1 || M == 2 || M == 3 {
		return []int{M}, nil
	}

	results := make([]int, 0)

	// (1) もしMが偶数なら、素因数2を出力する
	if M%2 == 0 {
		results = append(results, 2)
		return ShorFactorizationTwice(results, M/2)
	}

	// (2) M = a^b (a >= 1, b >= 2)かどうかを確かめる
	// もしそうであれば、素因数aを出力する
	exponent, err := num.CalcExponent(M)
	fmt.Println("CalcExponent", exponent, err)

	if err == nil {
		for i := 0; i < exponent.Exp; i++ {
			results = append(results, exponent.Base)
		}

		return results, nil
	}

	for count := 0; count < 30; count++ {
		fmt.Println("ShorFactorization loop start count=", count)

		randBigInt, err := num.RandomInt(int64(M - 2))
		if err != nil {
			return nil, err
		}

		// (3) 1 から M - 1 の間で任意に x を選ぶ。
		x := num.Int64ToInt(randBigInt.Int64() + 1)
		fmt.Println("ShorFactorization#after num.Int64ToInt", x)

		// もし x と M の最大公約数 (gcd(x, M)) が 1 より大きい (gcd(x, M) > 1)ならば、gcd(x, M)を出力する
		gcd := num.Gcd(x, M)
		fmt.Println("gcd", gcd)
		if gcd > 1 {
			return computedAfterGcd(results, gcd, M)
		}

		// (4) x, M (x < M) の位数 r を計算する (x^r mod M = 1)
		r := discoverClassicOrder(x, M)

		// (5) もし r が偶数であり、x^(r/2) mod M != 1 ならば、gcd(x^(r/2) - 1, M)とgcd(x^(r/2) + 1, M)を計算する
		// もしこれらのうち一つが M の因数なら、それを出力する。だめなら(3)へ戻る。
		if r%2 != 0 {
			fmt.Println("order discovery failure")
			continue
		}

		fac1 := num.Gcd(int(math.Pow(float64(x), float64(r/2)))-1, M)
		if fac1 != 1 && fac1 != M {
			return computedAfterGcd(results, fac1, M)
		}

		fac2 := num.Gcd(int(math.Pow(float64(x), float64(r/2)))+1, M)
		if fac2 != 1 && fac2 != M {
			return computedAfterGcd(results, fac2, M)
		}

		fmt.Println("factor estimation failure")
	}

	if M != 1 {
		results = append(results, M)
	}

	fmt.Println("[Failed] ShorFactorization results=", results)
	return results, nil
}

// divisionTwoAsFarAsPossible 可能な限り2で割り続ける 第一引数は2で割った回数、第二引数は割り切った結果
func divisionTwoAsFarAsPossible(M int) (int, int) {
	count := 0

	tmp := M
	for tmp%2 == 0 {
		count++
		tmp /= 2
	}

	return count, tmp
}

// discoverClassicOrder 古典の位数探索
func discoverClassicOrder(a, N int) int {
	fmt.Println("discoverClassicOrder args", a, N)

	i := 1
	for i = 1; i < N; i++ {
		x := int(math.Pow(float64(a), float64(i))) % N
		if x == 1 {
			fmt.Println("discoverClassicOrder retVal", i)
			return i
		}
	}

	fmt.Println("discoverClassicOrder retVal", i)
	return i
}

func ShorFactorizationTwice(baseResults []int, newM int) ([]int, error) {
	fmt.Println("ShorFactorizationTwice", baseResults, newM)

	newResults, err := ShorFactorization(newM)
	fmt.Println("ShorFactorizationTwice#after ShorFactorization", newResults, err)
	if err != nil {
		fmt.Println("computedAfterGcd ERROR", err)
	}

	results := make([]int, 0, len(baseResults)+len(newResults))

	for _, result := range baseResults {
		results = append(results, result)
	}
	for _, result := range newResults {
		results = append(results, result)
	}

	sort.Ints(results)
	return results, nil
}

func computedAfterGcd(baseResults []int, resultGcd int, M int) ([]int, error) {
	fmt.Println("computedAfterGcd", baseResults, resultGcd, M)

	newResults1, err := ShorFactorization(resultGcd)
	fmt.Println("computedAfterGcd#after ShorFactorization1", newResults1, err)
	if err != nil {
		fmt.Println("computedAfterGcd ERROR", err)
	}
	newResults2, err := ShorFactorization(M / resultGcd)
	fmt.Println("computedAfterGcd#after ShorFactorization2", newResults2, err)
	if err != nil {
		fmt.Println("computedAfterGcd ERROR", err)
	}

	results := make([]int, 0, len(baseResults)+len(newResults1)+len(newResults2))

	for _, result := range baseResults {
		results = append(results, result)
	}
	for _, result := range newResults1 {
		results = append(results, result)
	}
	for _, result := range newResults2 {
		results = append(results, result)
	}

	sort.Ints(results)
	fmt.Println("computedAfterGcd retVal", results)
	return results, nil
}
