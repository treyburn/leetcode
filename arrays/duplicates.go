package arrays

type intSet map[int]struct{}

func containsDuplicate(nums []int) bool {
	set := make(intSet, 0)
	var duplicates bool

	for _, n := range nums {
		if _, ok := set[n]; !ok {
			set[n] = struct{}{}
		} else {
			duplicates = true
		}
	}

	return duplicates
}
