package main;
import "fmt"

func cloneEvenNumbers(numbers []int) []int {
	seenIndex := len(numbers) -1 ;
	placingIndex := seenIndex;

	for i:= seenIndex; i >= 0; i-- {
		if (numbers[i] != -1) {
			numbers[placingIndex] = numbers[i]
			placingIndex -= 1
			if numbers[i] % 2 == 0 {
				numbers[placingIndex] = numbers[i]
				placingIndex--;
			}
		}
	}
	return numbers;
}

func main() {
	input := []int{1,2,3,4,5,9,-1,-1}
	result := cloneEvenNumbers(input)
	fmt.Println(result)

}