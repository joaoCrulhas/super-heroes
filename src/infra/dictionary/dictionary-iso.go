package dictionary

type DictionaryIsoAlphabetic struct {
	alphabet   string
	dictionary map[rune]int
}

func NewDictionaryIsoAlphabetic(alphabet string, dictionary map[rune]int) *DictionaryIsoAlphabetic {
	return &DictionaryIsoAlphabetic{
		alphabet:   alphabet,
		dictionary: Compute(alphabet),
	}
}

func (d *DictionaryIsoAlphabetic) GetKey(input rune) int {
	key := d.dictionary[input]
	return key
}

func (d *DictionaryIsoAlphabetic) GetValue(key int) rune {
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
