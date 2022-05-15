package reverse_string

func ReverseString(input string) (output string) {
	if len(input) < 2 {
		return input
	}
	outputByte := make([]byte, 0, len(input))
	start := len(input) - 1
	for i := start; i >= 0; i-- {
		outputByte[start-i] = input[i]
	}
	output = string(outputByte)
	return output
}
