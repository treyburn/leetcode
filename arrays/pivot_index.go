package arrays

func pivotIndex(nums []int) int {
	const notFound = -1
	var idx int = 0
	var leftSum int
	var rightSum = sum(nums)

	for idx < len(nums) {
		rightSum -= nums[idx]
		if idx > 0 {
			leftSum += nums[idx-1]
		}
		if leftSum == rightSum {
			return idx
		}
		idx++
	}

	return notFound
}

func sum(nums []int) int {
	var sum int

	for _, n := range nums {
		sum += n
	}

	return sum
}