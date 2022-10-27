package main

import "fmt"

func main() {

	var numbers = []int{5, 2, -10, 4, 8, 9, -20, 5, 20, 4, 8}

	maxNumber := Max(numbers)
	fmt.Println(maxNumber)

}

func Max(numbers []int) int {
	var max int
	for i, num := range numbers {
		if i == 0 {
			max = num
		}
		if num > max {
			max = num
		}
	}

	return max
}

func MaxIndex(numbers []int) int {
	var max int
	var index int
	for i, num := range numbers {
		if i == 0 {
			index = i
			max = num
		}
		if num > max {
			index = i
			max = num
		}
	}

	return index
}
