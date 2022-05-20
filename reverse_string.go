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
<<<<<<< HEAD

=======
>>>>>>> de8fc06f9fd06c34cdb175421e589b42ef40eb6d
	return output
}
