package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	inputBytes := []byte(input)
	length := len(inputBytes)
	for n := 0; n < length/2; n++ {
		//Reverser a string in bytes
		inputBytes[n], inputBytes[length-1-n] = inputBytes[length-1-n], inputBytes[n]
	}
	output = string(inputBytes)
	return output
}
