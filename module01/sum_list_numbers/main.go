package main

import "fmt"

func sum(numbers []float64) float64 {
	somat := 0.0
	for _, num := range numbers {
		somat += num
	}

	return somat
}

func main() {
	//numbers := []float64{10.0034, 20.01, 30.130, 1000.012, 2.999}
	numbers := []float64{}
	result := sum(numbers)
	fmt.Println(result)
}
