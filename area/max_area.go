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

// this is a test function for running benchmarks to see if the `if x := foo(); x != y {bar}` syntax is slower or faster
func originalMaxArea(height []int) int {
	var maxArea int
	left := 0
	right := len(height) - 1

	for left < right {
		if currentArea := min(height[right], height[left]) * (right - left); currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
