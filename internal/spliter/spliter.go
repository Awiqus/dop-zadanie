package spliter

import (
	"strings"
)

func Spliter(text string) ([]string, []string) {
	lines := strings.Split(text, "\n")
	words := make([]string, 0)

	for i := range lines {
		splittedLines := strings.Split(lines[i], " ")
		for _, k := range splittedLines {
			words = append(words, k)
		}
	}

	return words, lines
}
