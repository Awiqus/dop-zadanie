package frequency

func MostFrequentWords(words []string) (int, string) {
	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v] += 1
	}
	mostUsedWord := ""
	mostUsedCount := wordsMap[words[0]]
	for word, count := range wordsMap {
		if count > mostUsedCount {
			mostUsedCount = count
			mostUsedWord = word
		}
	}
	return mostUsedCount, mostUsedWord
}

func LetterFrequency(words []string) (int, string) {
	lettersMap := make(map[string]int)
	var letters []rune
	var runes []rune
	for _, v := range words {
		runes = []rune(v)
		letters = append(letters, runes...)
	}
	for _, v := range letters {
		_, ok := lettersMap[string(v)]
		if !ok {
			lettersMap[string(v)] = 1
			continue
		}
		lettersMap[string(v)] += 1
	}
	var mostFrequentLetterCount int
	var mostFrequentLetter string
	for letter := range lettersMap {
		if lettersMap[letter] > mostFrequentLetterCount {
			mostFrequentLetterCount = lettersMap[letter]
			mostFrequentLetter = letter
		}
	}

	return mostFrequentLetterCount, mostFrequentLetter
}
