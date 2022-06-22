package binary_search

func search(nums []int, target int) int {
	const notFound = -1
	var low, mid int
	var high = len(nums) - 1

	for low <= high {
		mid = (low + high) / 2
		found := nums[mid]
		switch {
		case found == target:
			return mid
		case found < target:
			low = mid + 1
		case found > target:
			high = mid - 1
		}
	}

	return notFound
}
