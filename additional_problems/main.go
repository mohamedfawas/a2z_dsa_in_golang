package main

import (
	probsreview "addprobs/probs_review"
	"fmt"
)

func main() {
	str := "Hello world"
	remChar := 'l'

	result := probsreview.RemoveChar(str, remChar)
	fmt.Println(result)
}
