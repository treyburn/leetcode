package substring

import "github.com/treyburn/generic"

// we use a sliding window to find the longest substring
func lengthOfLongestSubstring(s string) int {
	var longest int
	var chars = make(map[rune]int, 0)
	var right, left int // this determines our sliding window

	str := []rune(s)
	for right < len(str) { // keep going until end of string
		char := str[right]

		idx, ok := chars[char]
		if ok && idx >= left && idx < right { // by comparing idx to right and left we rule out letters where the window has already moved by
			left = idx + 1 // if the letter is a duplicate within the window then we shift the window bounds to the right by 1
		}

		longest = generic.Max(longest, right-left+1)

		chars[char] = right // if we have a duplicates then we will override that characters position to the most recently found
		right++
	}
	return longest
}
