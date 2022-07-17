package strings

var lookup = map[int][]rune{
	1:    []rune("I"),
	5:    []rune("V"),
	10:   []rune("X"),
	50:   []rune("L"),
	100:  []rune("C"),
	500:  []rune("D"),
	1000: []rune("M"),
}

func intToRoman(num int) string {
	iters := []int{1000, 500, 100, 50, 10, 5, 1}
	subtract := []int{100, 100, 100, 10, 10, 1, 1}
	output := make([]rune, 0)
	pos := 0

	for pos < len(iters) {
		if pos > 0 && num < iters[pos-1] && num >= iters[pos-1]-subtract[pos] {
			sub := subtract[pos]
			lead := iters[pos-1]
			output = append(output, lookup[sub]...)
			output = append(output, lookup[lead]...)
			num -= lead - sub
			continue
		}
		if num >= iters[pos] {
			lead := iters[pos]
			output = append(output, lookup[lead]...)
			num -= lead
		} else {
			pos++
			if pos == len(iters) {
				break
			}
		}
	}

	return string(output)
}
