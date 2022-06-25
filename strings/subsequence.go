package strings

func isSubsequence(s string, t string) bool {
	var sRune = []rune(s)
	var position int

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	for _, r := range []rune(t) {
		if r == sRune[position] {
			position++
		}
		if position == len(sRune) {
			return true
		}
	}
	return false
}

// attempt 2 after learning greedy algo
// interesting enough - my original solution did better than the "optimized" solution in terms of time complexity
// and identical in terms of space complexity
// this may be due to the test data not being sufficiently complex. granted - the solutions are pretty close to
// identical in logic - just structured slightly differently
func isSubsequence2(s string, t string) bool {
	sRunes := []rune(s)
	tRunes := []rune(t)
	leftBound := len(s)
	rightBound := len(t)
	var leftIndex, rightIndex int

	for {
		if leftIndex == leftBound {
			return true
		}
		if rightIndex == rightBound {
			return false
		}
		if sRunes[leftIndex] == tRunes[rightIndex] {
			leftIndex++
		}
		rightIndex++
	}
}
