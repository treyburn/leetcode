package area

// we start at both sides and move inwards, calculating area each step of the way
// we determine which direction to move inwards based on which is shorter
// this is an O(n) solution with O(1) space complexity
func maxArea(height []int) int {
	var maxArea int
	left := 0
	right := len(height) - 1

	for left < right {
		currentArea := min(height[right], height[left]) * (right - left)
		maxArea = max(currentArea, maxArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
