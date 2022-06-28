package add_nums

func twoSum(nums []int, target int) []int {
	var positions []int
	var set = make(map[int]int, 0)

	for idx, val := range nums {
		want := target - val
		pos, ok := set[want]
		if ok {
			positions = []int{pos, idx}
			return positions
		}
		set[val] = idx
	}
	return positions
}
