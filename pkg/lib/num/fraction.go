package num

// ContinuedFractionExpansion a/bを連分数展開
func ContinuedFractionExpansion(a, b float64, loop int) []float64 {
	results := make([]float64, 0, loop)
	nextA := a
	nextB := b

	for i := 0; i < loop; i++ {
		cA := nextA
		cB := nextB

		if cB == 0 {
			break
		}

		currentNum := RoundDown(cA/cB, 0)
		results = append(results, currentNum)

		tempNextA := cA - currentNum*cB

		nextA = cB
		nextB = tempNextA
	}

	return results
}
