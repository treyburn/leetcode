package arrays

func runningSum(nums []int) []int {
	var output = make([]int, len(nums))
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		output[i] = sum
	}

	return output
}