package reverse_string

func ReverseString(input string) (output string) {
	var runes = []rune(input)
	var length = len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	output = string(runes)

	return output
}
