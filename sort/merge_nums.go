package sort

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ptr1 = m - 1 // start at the end of real values in nums1
	var ptr2 = n - 1 // start at the end of nums2

	for i := m + n - 1; i >= 0; i-- { // start at the end of nums1 (all 0's) and move backwards
		if ptr2 < 0 { // return early if we run out of nums2 - as the rest of nums1 will have been smaller
			return
		}
		// write the smallest value to the current position of i, then move backwards
		if ptr1 >= 0 && nums1[ptr1] > nums2[ptr2] {
			nums1[i] = nums1[ptr1]
			ptr1--
		} else {
			nums1[i] = nums2[ptr2]
			ptr2--
		}
	}
}
