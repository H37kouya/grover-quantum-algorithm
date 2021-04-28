package service

import (
	"fmt"
	"grover-quantum-search/pkg/domain/valueObject"
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
			return doubleCheckShorFactorization(results, gcd, M)
		}

		// (4) x, M (x < M) の位数 r を計算する (x^r mod M = 1)
		r, err := discoverQuantumOrder(x, M)
		if err != nil {
			return nil, err
		}

		// (5) もし r が偶数であり、x^(r/2) mod M != 1 ならば、gcd(x^(r/2) - 1, M)とgcd(x^(r/2) + 1, M)を計算する
		// もしこれらのうち一つが M の因数なら、それを出力する。だめなら(3)へ戻る。
		if r%2 != 0 {
			fmt.Println("order discovery failure")
			continue
		}

		fac1 := num.Gcd(int(math.Pow(float64(x), float64(r/2)))-1, M)
		if fac1 != 1 && fac1 != M {
			return doubleCheckShorFactorization(results, fac1, M)
		}

		fac2 := num.Gcd(int(math.Pow(float64(x), float64(r/2)))+1, M)
		if fac2 != 1 && fac2 != M {
			return doubleCheckShorFactorization(results, fac2, M)
		}

		fmt.Println("factor estimation failure")
	}

	if M != 1 {
		results = append(results, M)
	}

	fmt.Println("[Failed] ShorFactorization results=", results)
	return results, nil
}

// discoverQuantumOrder 量子の位数探索
func discoverQuantumOrder(x, M int) (int, error) {
	fmt.Println("discoverQuantumOrder args", x, M)

	// 1) 整数 N と互いに素になる整数 x を選ぶ。
	// 2) 初期状態を準備する。
	//    第1レジスタ：s / r の位相推定結果を必要な精度で納めるため t 量子ビット (|0> に初期化)
	//    第2レジスタ：N を入力する計算用の L 量子ビット (|1> に初期化)
	// 3) 第1レジスタすべてにアダマールゲートを作用する。
	// 4) 制御ユニタリゲート U(x, M) を作用させる。
	// 5) 第1レジスタに量子フーリエ逆変換を行う。
	// 6) 第1レジスタを測定し s / r を得る。
	// 7) 連分数アルゴリズムを適用し位数 r を決定する。

	// MEMO
	// t = 2 L + 1 + log(3 + 1 / 2ε)
	// ε は r を推定する手続きで失敗する確率の上限

	// M の 2進数表記の桁数
	n, err := valueObject.NewNByTenDecimalNumber(M)
	if err != nil {
		return 0, err
	}
	// 位相推定の精度 t = 2 L + 1 + log(3 + 1 / 2ε)
	t, err := valueObject.NewN(2*n.Get() + 1 + 2)
	if err != nil {
		return 0, err
	}

	fmt.Println("discoverQuantumOrder M のビット数", n)
	fmt.Println("discoverQuantumOrder t = 2 L + 1 + log(3 + 1 / 2ε)", t)

	// 4) 制御ユニタリゲート U(x, N) を作用させる。
	rArr := ControlUnitaryGate(x, M, t)
	r := len(rArr) // 周期

	/*	qubits := collection.MakeNQubits(t)
		for i := 0; i < t.ElementCount(); i++ {
			qubits = append(qubits, math.Exp(2 * math.Pi *))
		}*/

	fmt.Println("4) 制御ユニタリゲート U(x, N) を作用させる", r, rArr)

	fmt.Println("discoverQuantumOrder retVal", 0)
	return 0, nil
}

// ControlUnitaryGate 制御ユニタリゲート U(x, M) nはビット数
func ControlUnitaryGate(x, M int, n valueObject.N) []int {
	fmt.Println("ControlUnitaryGate", x, M, n)

	rArr := make([]int, 0, n.ElementCount()) // 周期の配列
	for k := 0; k <= n.ElementCount(); k++ {
		z := int(math.Pow(float64(x), float64(k))) % M

		for _, rVal := range rArr {
			if rVal == z {
				sort.Ints(rArr)
				fmt.Println("ControlUnitaryGate retVal", rArr)
				return rArr
			}
		}

		rArr = append(rArr, z)
	}

	sort.Ints(rArr)
	fmt.Println("ControlUnitaryGate retVal", rArr)
	return rArr
}

// ShorFactorizationTwice 因数が素因数分解ができるかどうかを確認する
func ShorFactorizationTwice(baseResults []int, newM int) ([]int, error) {
	fmt.Println("ShorFactorizationTwice", baseResults, newM)

	newResults, err := ShorFactorization(newM)
	fmt.Println("ShorFactorizationTwice#after ShorFactorization", newResults, err)
	if err != nil {
		fmt.Println("doubleCheckShorFactorization ERROR", err)
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

// doubleCheckShorFactorization 2つの因数をさらに素因数分解できるかを確認する
func doubleCheckShorFactorization(baseResults []int, resultGcd int, M int) ([]int, error) {
	fmt.Println("doubleCheckShorFactorization", baseResults, resultGcd, M)

	newResults1, err := ShorFactorization(resultGcd)
	fmt.Println("doubleCheckShorFactorization#after ShorFactorization1", newResults1, err)
	if err != nil {
		fmt.Println("doubleCheckShorFactorization ERROR", err)
	}
	newResults2, err := ShorFactorization(M / resultGcd)
	fmt.Println("doubleCheckShorFactorization#after ShorFactorization2", newResults2, err)
	if err != nil {
		fmt.Println("doubleCheckShorFactorization ERROR", err)
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
	fmt.Println("doubleCheckShorFactorization retVal", results)
	return results, nil
}
