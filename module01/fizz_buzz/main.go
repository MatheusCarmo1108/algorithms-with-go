package main

import "fmt"

func FizzBuzz(num_max int) {
	for i := 1; i <= num_max; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		if i != num_max {
			fmt.Print(", ")
		}
	}
}

func main() {
	number := 15
	FizzBuzz(number)
}
