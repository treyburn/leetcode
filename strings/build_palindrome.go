package strings

func longestPalindrome(s string) int {
	var set = make(map[rune]int)
	var usedFreebie bool
	var length int

	for _, char := range []rune(s) {
		val, ok := set[char]
		if !ok {
			set[char] = 1
		} else {
			set[char] = val + 1
		}
	}

	for _, val := range set {
		switch {
		case val > 1:
			remainder := val % 2
			if remainder == 1 && !usedFreebie {
				usedFreebie = true
				remainder--
			}
			length += val - remainder
		default:
			if !usedFreebie {
				usedFreebie = true
				length++
			}
		}

	}
	return length
}
