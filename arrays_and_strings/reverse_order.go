package main

import "fmt"

func reverse(numbers []int) []int {
	startI := 0
	endI := len(numbers) - 1
	for startI < endI {
		tmp := numbers[startI]
		numbers[startI] = numbers[endI]
		numbers[endI] = tmp
		startI++
		endI--
	}
	return numbers
}

func main() {
	result := reverse([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
}
