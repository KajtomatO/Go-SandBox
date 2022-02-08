package iteration

// This function wil take strink and integer and return integer repeated integer times
func Repeat(word string, rep int) (repeatedString string) {
	for i := 0; i < rep; i++ {
		repeatedString += word
	}
	return
}
