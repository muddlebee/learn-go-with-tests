package iteration

const repeatCount = 5

// Repeat returns character repeated 5 times.
func Repeat(character string) string {
	var finalStr string
	for i := 0; i < repeatCount; i++ {
		finalStr += character
	}
	return finalStr
}
