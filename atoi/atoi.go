package atoi

var convMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

const (
	negative = "-"
	positive = "+"
	max      = 2147483647
	min      = -2147483648
)

func myAtoi(s string) int {
	var output int
	var nums = make([]int, 0)
	var isNegative bool
	var modified bool

	str := []rune(s)

loop:
	for _, aNum := range str {
		if iNum, ok := convMap[string(aNum)]; ok {
			nums = append(nums, iNum)
		} else {
			switch string(aNum) {
			case negative:
				if modified || len(nums) > 0 {
					break loop
				}
				isNegative = true
				modified = true
			case positive:
				if modified || len(nums) > 0 {
					break loop
				}
				modified = true
			case " ":
				if modified || len(nums) > 0 {
					break loop
				}
			default:
				break loop
			}
		}

	}

	nums = trimLeadingZeros(nums)

	if len(nums) > 10 {
		if isNegative {
			return min
		} else {
			return max
		}
	}

	for _, iNum := range nums {
		output = output*10 + iNum
	}

	if isNegative {
		output = output * -1
	}

	if output > max {
		return max
	}

	if output < min {
		return min
	}

	return output
}

func trimLeadingZeros(n []int) []int {
	trimming := true
	out := make([]int, 0)
	for _, v := range n {
		if v != 0 || trimming == false {
			trimming = false
			out = append(out, v)
		}
	}
	return out
}
