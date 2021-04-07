package main

import "fmt"

func NumInList(list []int, num int) bool {
	for _, item := range list {
		if item == num {
			return true
		}
	}
	return false
}

func main() {
	lista := []int{10, 20, 30, 40, 50, 60, 70, 80}
	num_certo := 10
	fmt.Println(NumInList(lista, num_certo))

}
