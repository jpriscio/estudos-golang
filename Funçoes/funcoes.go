package main

import "fmt"

func val() (int, int) {
	return 3, 5
}

func somas(soma ...int) int {
	numeros := 0

	for _, total := range soma {
		numeros += total
	}

	return numeros
}

func main() {

	slice := []int{10, 10, 10, 50}

	fmt.Println(somas(slice...))

}
