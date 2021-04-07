package main

import "fmt"

func reverse(frase string) string {
	res := ""
	for _, r := range frase {
		res = string(r) + res
	}
	return res
}

func main() {
	name := "Matheus do Carmo de Oliveira Neto"
	fmt.Println(reverse(name))
}
