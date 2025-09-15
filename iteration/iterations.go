package iteration

const repeatCount = 5

func Repeat(character rune) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += string(character)
	}
	return repeated
}