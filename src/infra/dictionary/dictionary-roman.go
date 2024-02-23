package dictionary

type DictionaryRomanAlphabetic struct {
	alphabet   string
	dictionary map[rune]int
}

func NewDictionaryRomanAlphabetic(alphabet string, dictionary map[rune]int) *DictionaryRomanAlphabetic {
	return &DictionaryRomanAlphabetic{
		alphabet:   alphabet,
		dictionary: Compute(alphabet),
	}
}

func (d *DictionaryRomanAlphabetic) GetKey(input rune) int {
	key := d.dictionary[input]
	return key
}

func (d *DictionaryRomanAlphabetic) GetValue(key int) rune {
	for k, v := range d.dictionary {
		if v == key {
			return k
		}
	}
	return 0
}

func Compute(alphabet string) map[rune]int {
	alphabetMap := make(map[rune]int)
	for i, letter := range alphabet {
		alphabetMap[letter] = i + 1
	}
	return alphabetMap
}
