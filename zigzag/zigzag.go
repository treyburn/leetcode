package zigzag

func convert(s string, numRows int) string {
	var output string
	var cursor, wordIndex int // cursor corresponds to
	var reverse bool          // this controls which direction the cursor moves
	repo := make([][]rune, numRows)

	str := []rune(s)

	for i := 0; i < len(str); i++ {
		repo[cursor] = append(repo[cursor], str[wordIndex])
		wordIndex++
		if numRows > 1 { // we only need to move the cursor if there's more than 1 row
			switch cursor {
			case 0:
				reverse = false
			case numRows - 1:
				reverse = true
			}
			switch reverse {
			case true:
				cursor--
			case false:
				cursor++
			}
		}
	}

	for _, substr := range repo {
		output += string(substr)
	}

	return output
}
