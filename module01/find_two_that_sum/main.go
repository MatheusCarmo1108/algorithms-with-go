package main

import "fmt"

func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, num1 := range numbers {
		for j, num2 := range numbers {
			if i == j {
				continue
			}
			if num1+num2 == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	slice := []int{1, 2, 3, 4}
	somat := 7

	fmt.Println(FindTwoThatSum(slice, somat))

}
