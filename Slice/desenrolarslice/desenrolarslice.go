package main

import "fmt"

func desenrolarslice(numeros ...int) int {
	valor := 0

	for _, n := range numeros {
		valor = valor + n

	}
	return valor

}

func main() {

	slice := []int{10, 10, 10, 100}

	fmt.Println(desenrolarslice(slice...))

}
