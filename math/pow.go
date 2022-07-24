package math

func myPow(x float64, n int) float64 {
	answer := 1.00000
	if n == 0 {
		return answer
	}

	origN := n
	n = abs(n)

	for n > 0 {
		lastBit := n & 1

		if lastBit == 1 {
			answer *= x
		}

		x *= x

		n >>= 1
	}

	if origN < 0 {
		return trim(1 / answer)
	}

	return trim(answer)
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func trim(x float64) float64 {
	const n = 100000
	return float64(int64(x*n)) / n
}
