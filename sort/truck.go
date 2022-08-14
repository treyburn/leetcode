package sort

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var count int
	var inTruck int

	sortedBoxes := sortBoxes(boxTypes)

	for _, box := range sortedBoxes {
		if inTruck == truckSize {
			return count
		}
		if box[0] <= (truckSize - inTruck) {
			count += box[0] * box[1]
			inTruck += box[0]
		} else {
			count += (truckSize - inTruck) * box[1]
			return count
		}
	}
	return count
}

func sortBoxes(boxes [][]int) [][]int {
	if len(boxes) < 2 {
		return boxes
	}

	var lesser [][]int
	var greater [][]int

	pivot := boxes[0]

	for _, box := range boxes[1:] {
		if box[1] >= pivot[1] {
			greater = append(greater, box)
		} else {
			lesser = append(lesser, box)
		}
	}

	return append(append(sortBoxes(greater), pivot), sortBoxes(lesser)...)
}
