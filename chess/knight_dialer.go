package chess

var dialOpts = map[int][]int{
	0: {4, 6},
	1: {6, 8},
	2: {7, 9},
	3: {4, 8},
	4: {0, 3, 9},
	5: {},
	6: {0, 1, 7},
	7: {2, 6},
	8: {1, 3},
	9: {2, 4},
}

type pair struct {
	x int
	n int
}

var memo = map[pair]int{
	{0, 1}: 1,
	{1, 1}: 1,
	{2, 1}: 1,
	{3, 1}: 1,
	{4, 1}: 1,
	{5, 1}: 1,
	{6, 1}: 1,
	{7, 1}: 1,
	{8, 1}: 1,
	{9, 1}: 1,
}

const mod = 1e9 + 7

func knightDialer(n int) int {
	var total int
	for i := 0; i < 10; i++ {
		if i == 5 {
			if n == 1 {
				total++
			}
			continue
		}
		total += solve(i, n)
	}

	return total % mod
}

func solve(pos int, n int) int {
	if v, ok := memo[pair{pos, n}]; ok {
		return v
	}

	var answer int
	for _, num := range dialOpts[pos] {
		answer += solve(num, n-1)
	}
	memo[pair{pos, n}] = answer % mod
	return answer % mod
}
