package reverse_string

func ReverseString(input string) (output string) {
	if len(input) < 2 {
		return input
	}
	outputRune := make([]rune, 0, len(input))
	inputRune := []rune(input)
	start := len(input) - 1
	for i := start; i >= 0; i-- {
		outputRune = append(outputRune, inputRune[i])
	}
	return string(outputRune)
}
