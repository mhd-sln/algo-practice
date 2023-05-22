package main

import (
	"fmt"
)

func getReverse(sentence string) string {
	reversed := ""

	wordEndIndex := len(sentence)
	sentenceIndex := len(sentence) - 1
	for ; sentenceIndex > 0; sentenceIndex-- {
		if sentence[sentenceIndex] == ' ' {
			word := sentence[sentenceIndex+1 : wordEndIndex]
			reversed += word
			reversed += " "
			wordEndIndex = sentenceIndex
		}
	}
	reversed += sentence[sentenceIndex: wordEndIndex ]
	return reversed
}
func main() {
	sentence := "I live in a house nearby"
	reverse := getReverse(sentence)
	fmt.Println(reverse)
}
