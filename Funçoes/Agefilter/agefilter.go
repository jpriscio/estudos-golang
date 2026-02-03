package main

import "fmt"

func createfilter(min, max int) func(age int) (bool, int) {
	contador := 0
	return func(age int) (bool, int) {
		v := age >= min && age <= max
		if v == false {
			contador++
		}

		return v, contador
	}

}

func main() {

	ageFilter := createfilter(18, 28)

	fmt.Println(ageFilter(23))
	fmt.Println(ageFilter(17))
	fmt.Println(ageFilter(15))

}
