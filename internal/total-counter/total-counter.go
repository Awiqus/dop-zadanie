package totalcounter

func TotalWords(words []string) int {
	var counter int
	for _, v := range words {
		if v == "," || v == "" {
			continue
		}
		counter++
	}
	return counter
}

func TotalLines(lines []string) int {
	return len(lines)
}
