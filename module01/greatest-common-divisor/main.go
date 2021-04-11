package main

import "fmt"

func main() {
	fmt.Println(GCD(100, 9))
}

func GCD(num_1, num_2 int) int {
	for {
		if num_2 == 0 {
			return num_1
		}

		num_1 = num_2
		num_2 = num_1 % num_2
	}
}
