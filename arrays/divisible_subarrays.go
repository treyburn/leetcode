package arrays

// brute force approach - not efficient enough
func subarraysDivByK_orig(nums []int, k int) int {
	var count int
	max := len(nums)
	var left, right int

	for left < max {
		current := 0
		right = left
		for right < max {
			current += nums[right]
			if current%k == 0 {
				count++
			}
			right++
		}
		left++
	}

	return count
}

func subarraysDivByK(nums []int, k int) int {
	var count int
	mod := make([]int, k+1)

	cumulativeSum := 0
	for i := 0; i < len(nums); i++ {
		cumulativeSum += nums[i]

		mod[((cumulativeSum%k)+k)%k]++
	}

	for j := 0; j < k; j++ {
		if mod[j] > 1 {
			count += (mod[j] * (mod[j] - 1)) / 2
		}
	}

	count += mod[0]

	return count
}
