package main

func percentageLetter(s string, letter byte) int {
	totalLength := len(s)
	count := 0
	for i := 0; i < totalLength; i++ {
		if s[i] == letter {
			count++
		}
	}

	perc := (count * 100) / totalLength
	return perc
}
