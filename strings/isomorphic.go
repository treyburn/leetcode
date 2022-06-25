package strings

type isomorphicHash struct {
	tally     int
	positions map[rune][]int
	chars     map[int]rune
}

func newIsomorphicHash() isomorphicHash {
	c := make(map[rune][]int, 0)
	p := make(map[int]rune, 0)

	return isomorphicHash{
		tally:     0,
		positions: c,
		chars:     p,
	}
}

func isIsomorphic(s string, t string) bool {
	const isomorphic = true

	sHash := hash([]rune(s))
	tHash := hash([]rune(t))

	if sHash.tally != tHash.tally {
		return !isomorphic
	}

	for i := 0; i < sHash.tally; i++ {
		sChar := sHash.chars[i]
		tChar := tHash.chars[i]

		sPositions := sHash.positions[sChar]
		tPositions := tHash.positions[tChar]

		if len(sPositions) != len(tPositions) {
			return !isomorphic
		}

		for idx, val := range sPositions {
			if val != tPositions[idx] {
				return !isomorphic
			}
		}

	}

	return isomorphic
}

func hash(s []rune) isomorphicHash {
	h := newIsomorphicHash()

	for idx, char := range s {
		val, ok := h.positions[char]
		if !ok {
			h.tally++
			h.chars[idx] = char
		}
		h.positions[char] = append(val, idx)
	}

	return h
}

// Second attempt after reading how the algo is supposed to work
// results from the benchmark are that its 4 times faster than 1
// uses 1/5 of the memory - and has 2 allocs/op vs 45 allocs/op
func isIsomorphic2(s string, t string) bool {
	const isomorphic = true
	var forwardHash = make(map[rune]rune)
	var backwardHash = make(map[rune]rune)

	sRunes := []rune(s)
	tRunes := []rune(t)

	if len(sRunes) != len(tRunes) {
		return !isomorphic
	}

	for i := 0; i < len(sRunes); i++ {
		val1, ok1 := forwardHash[sRunes[i]]
		val2, ok2 := backwardHash[tRunes[i]]
		switch {
		case !ok1 && !ok2:
			forwardHash[sRunes[i]] = tRunes[i]
			backwardHash[tRunes[i]] = sRunes[i]
		case ok1 && ok2:
			if val1 != tRunes[i] || val2 != sRunes[i] {
				return !isomorphic
			}
		default:
			return !isomorphic
		}
	}

	return isomorphic
}
