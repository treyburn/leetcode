package substring

// were going to use a sliding window to find the longest palindrome
func longestPalindrome(s string) string {
	var longest []rune
	var right, left int // this determines our sliding window

	str := []rune(s)
	for left < len(str) {
		current := make([]rune, 0)
		for right < len(str) {
			current = append(current, str[right])
			if isPalindrome(current) && len(current) > len(longest) {
				longest = current
			}
			right++
		}
		left++
		right = left
	}

	return string(longest)
}

func isPalindrome(str []rune) bool {
	l := len(str)
	if l < 1 {
		return false
	}

	start := 0
	end := len(str) - 1

	for start <= end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true

}
