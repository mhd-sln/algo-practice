package main

import "fmt"

func getEvenNumberTwice(numbers []int) []int {
	lastNumberIndex := getLastNumberIndex(numbers)
	endIndex := len(numbers) - 1
	for ; lastNumberIndex >= 0; lastNumberIndex-- {
		numbers[endIndex] = numbers[lastNumberIndex]
		endIndex--
		if numbers[lastNumberIndex]%2 == 0 {
			numbers[endIndex] = numbers[lastNumberIndex]
			endIndex--
		}
	}
	return numbers
}

func getLastNumberIndex(numbers []int) int {
	i := len(numbers) - 1
	for ; numbers[i] == -1; i-- {
	}
	return i
}

func main() {
	numbers := []int{1, 2, 3, 4, 6, 8, -1, -1, -1, -1}
	evenNumberTwice := getEvenNumberTwice(numbers)
	fmt.Println(evenNumberTwice)
}
