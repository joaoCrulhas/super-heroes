package domain

type (
	Dictionary interface {
		GetKey(input rune) int
		GetValue(key int) rune
		GetAlphabetLength() int
	}
)
