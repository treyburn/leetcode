package binary_search

func searchInsert(nums []int, target int) int {
	var low, mid int
	var high = len(nums) - 1
	var found int

	for low <= high {
		mid = (low + high) / 2
		found = nums[mid]
		switch {
		case found == target:
			return mid
		case found < target:
			low = mid + 1
		case found > target:
			high = mid - 1
		}
	}

	if found < target {
		return mid + 1
	} else {
		return mid
	}
}
