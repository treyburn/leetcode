package arrays

// find the median value of 2 sorted arrays
// we are going to use a merge sort
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := merge(nums1, nums2)
	return median(merged)
}

func merge(arr1 []int, arr2 []int) []int {
	n1 := len(arr1)
	n2 := len(arr2)
	arr3 := make([]int, n1+n2)
	var arr1idx, arr2idx, arr3idx int

	// traverse both arrays
	for arr1idx < n1 && arr2idx < n2 {
		// Check if current element of first array is smaller than current element of second array.
		// If yes, store first array element and increment first array index.
		//Otherwise, do same with second array
		switch arr1[arr1idx] < arr2[arr2idx] {
		case true:
			arr3[arr3idx] = arr1[arr1idx]
			arr3idx++
			arr1idx++
		case false:
			arr3[arr3idx] = arr2[arr2idx]
			arr3idx++
			arr2idx++
		}
	}

	// drain remaining elements of arr1 - if any
	for arr1idx < n1 {
		arr3[arr3idx] = arr1[arr1idx]
		arr3idx++
		arr1idx++
	}

	//drain remaining elements of arr2 - if any
	for arr2idx < n2 {
		arr3[arr3idx] = arr2[arr2idx]
		arr3idx++
		arr2idx++
	}

	return arr3
}

func median(arr []int) float64 {
	l := len(arr)
	if l < 1 {
		panic("cannot get median of empty array")
	}

	switch l % 2 {
	default: // array has an odd number of elements
		return float64(arr[l/2])
	case 0: // array has an even number of elements
		floor := arr[l/2]
		ceil := arr[l/2-1]
		return float64(floor+ceil) / 2
	}
}
