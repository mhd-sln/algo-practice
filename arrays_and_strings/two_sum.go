package main

import "fmt"

func find_nums_for_sum(numbers []int, sum int) []int {
	startI := 0
	endI := len(numbers) -1
	for startI < endI {
		start := numbers[startI]
		end := numbers[endI]
		if start+end == sum {
			return []int{start, end}
		} else if start+end > sum {
			endI--
		} else {
			startI++
		}
	}
	return []int{}
}

func main() {
	input := []int{1, 2, 3, 4, 7, 9}
	sum := 7
	nums := find_nums_for_sum(input, sum)
	fmt.Println(nums)
}
