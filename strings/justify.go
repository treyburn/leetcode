package strings

func fullJustify(words []string, maxWidth int) []string {
	output := make([]string, 0)
	currentWords := make([]string, 0)
	var currentLen int

	for _, word := range words {
		tempLen := currentLen
		if len(currentWords) >= 1 {
			tempLen++
		}
		tempLen += len(word)
		if tempLen <= maxWidth {
			currentLen = tempLen
			currentWords = append(currentWords, word)
		} else {
			output = append(output, justify(currentWords, currentLen, maxWidth))
			currentWords = []string{word}
			currentLen = len(word)
		}
	}

	var finalRow string
	for _, word := range currentWords {
		if finalRow != "" {
			finalRow += " "
		}
		finalRow += word
	}

	if finalRow != "" {
		for len(finalRow) < maxWidth {
			finalRow += " "
		}
		output = append(output, finalRow)
	}

	return output
}

func justify(words []string, currentLen, width int) string {
	var output string
	var sep string
	numSeps := len(words) - 1
	var overflowSep int
	amtToSep := width - currentLen + numSeps
	if numSeps > 0 {
		sepSize := amtToSep / numSeps
		overflowSep = amtToSep % numSeps
		for i := 0; i < sepSize; i++ {
			sep += " "
		}
	}

	for _, word := range words {
		if output != "" {
			output += sep
			if overflowSep > 0 && numSeps > 0 {
				output += " "
				overflowSep--
			}
		}
		output += word
	}

	for len(output) < width {
		output += " "
	}

	return output
}
