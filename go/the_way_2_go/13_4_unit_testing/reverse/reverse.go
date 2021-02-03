package reverse

// reversing string
func Reverse(inputString string) string {
	reversedString := make([]rune, len(inputString))
	for index, char := range inputString {
		reversedString[len(reversedString)-index-1] = rune(char)
	}
	return string(reversedString)
}
