package main

import (
	"dop-zadanie/internal/frequency"
	"dop-zadanie/internal/spliter"
	totalcounter "dop-zadanie/internal/total-counter"
	"fmt"
)

func main() {
	text := "What a , wonderful DAY ! \n Goo d job a a a"

	words, lines := spliter.Spliter(text)
	tl := totalcounter.TotalLines(lines)
	tw := totalcounter.TotalWords(words)
	mfwc, mfw := frequency.MostFrequentWords(words)
	lfc, lf := frequency.LetterFrequency(words)

	fmt.Printf("Total words: %d\nTotal strokes: %d\nMost Frequenct Word: %s with %d use cases\nMost frequent letter: %s with uses %d\n", tw, tl, mfw, mfwc, lf, lfc)
}
