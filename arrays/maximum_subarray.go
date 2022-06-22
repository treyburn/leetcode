package arrays

// returns the value of the great sum of any contiguous subarray
// going to achieve O(n) by greedily scanning forward and resetting the current window of sums whenever the current value
// is greater than the same of the previous values in the window
func maxSubArray(nums []int) int {
	var max = nums[0]
	var current = nums[0]

	for i := 1; i < len(nums); i++ {
		var n = nums[i]
		current = maximum(n, current+n)
		max = maximum(max, current)
	}

	return max
}

func maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
