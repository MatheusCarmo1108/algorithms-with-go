package main

import "fmt"

func main() {
	number := 5
	fmt.Println(" =", Factorial(number))
}

func Factorial(num int) int {
	res := 1
	fmt.Printf("%v! = ", num)
	for i := num; i >= 1; i-- {
		fmt.Printf("%v", i)
		res = res * i
		if i != 1 {
			fmt.Printf("*")
		}
	}
	return res
}
