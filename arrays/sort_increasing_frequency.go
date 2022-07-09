package arrays

func frequencySort(nums []int) []int {
	var result = make([]int, len(nums))
	var frequency = make(map[int]int)
	var counts = make(map[int][]int)
	var distinctCounts []int

	// store frequency in hash map
	for _, n := range nums {
		val, ok := frequency[n]
		if !ok {
			frequency[n] = 1
		} else {
			frequency[n] = val + 1
		}
	}

	// create hashmap of count values and add to distinct counts array
	for k, v := range frequency {
		val, ok := counts[v]
		if !ok {
			counts[v] = []int{k}
			distinctCounts = append(distinctCounts, v)
		} else {
			// sort array of values first
			counts[v] = descendingQuickSort(append(val, k))
		}
	}

	// sort distinct counts array with quicksort
	distinctCounts = quickSort(distinctCounts)

	// traverse distinct counts and write to result
	pos := 0
	for _, count := range distinctCounts {
		vals := counts[count]
		for _, val := range vals {
			for i := 0; i < frequency[val]; i++ {
				result[pos] = val
				pos++
			}
		}
	}

	return result
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var lesser []int
	var greater []int

	pivot := arr[0]
	for _, val := range arr[1:] {
		if val >= pivot {
			greater = append(greater, val)
		} else {
			lesser = append(lesser, val)
		}
	}

	return append(append(quickSort(lesser), pivot), quickSort(greater)...)
}

func descendingQuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var lesser []int
	var greater []int

	pivot := arr[0]
	for _, val := range arr[1:] {
		if val >= pivot {
			greater = append(greater, val)
		} else {
			lesser = append(lesser, val)
		}
	}

	return append(append(descendingQuickSort(greater), pivot), descendingQuickSort(lesser)...)
}
