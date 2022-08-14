package dfs

var letters = map[rune][]string{
	[]rune("2")[0]: {"a", "b", "c"},
	[]rune("3")[0]: {"d", "e", "f"},
	[]rune("4")[0]: {"g", "h", "i"},
	[]rune("5")[0]: {"j", "k", "l"},
	[]rune("6")[0]: {"m", "n", "o"},
	[]rune("7")[0]: {"p", "q", "r", "s"},
	[]rune("8")[0]: {"t", "u", "v"},
	[]rune("9")[0]: {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	rDigits := []rune(digits)
	if len(rDigits) == 0 {
		return make([]string, 0)
	}

	return dfs(rDigits, "")
}

func dfs(nums []rune, word string) []string {
	if len(nums) == 0 {
		return []string{word}
	}

	results := make([]string, 0)
	for _, letter := range letters[nums[0]] {
		results = append(results, dfs(nums[1:], word+letter)...)
	}
	return results
}
