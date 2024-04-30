package main

import (
	"fmt"
	"strings"
)

func strHash(sentence string) map[string]int {
	sentence = strings.ReplaceAll(sentence, " ", "")
	sentence = strings.ToLower(sentence)
	str_arr := strings.Split(sentence, "")
	frq_map := make(map[string]int)
	for _, v := range str_arr {
		fmt.Println(v)
		frq_map[v]++
	}
	return frq_map
}

func main() {
	fmt.Println(strHash("Hai I am tony, I am here"))
}
