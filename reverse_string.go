package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	inputRunes := []rune(input)
	length := len(inputRunes)
	for n := 0; n < length/2; n++ {
		//Reverser a string in bytes
		inputRunes[n], inputRunes[length-1-n] = inputRunes[length-1-n], inputRunes[n]
	}
	output = string(inputRunes)
	return output
}
