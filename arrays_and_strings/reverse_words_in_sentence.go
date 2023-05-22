package main

import (
	"fmt"
)

func getReverse(sentence string) string {
	reversed := ""
	wordStartIndex := len(sentence) -1
	wordEndIndex := len(sentence)

	for ; wordStartIndex > 0; wordStartIndex-- {
		if(sentence[wordStartIndex] == ' ') {
			reversed += sentence[wordStartIndex+1:wordEndIndex]
			reversed += " "
			wordEndIndex = wordStartIndex
		}
	}
	reversed += sentence[0: wordEndIndex]
	return reversed
}
func main() {
	sentence := "I live in a house nearby"
	reverse := getReverse(sentence)
	fmt.Println(reverse)
}
