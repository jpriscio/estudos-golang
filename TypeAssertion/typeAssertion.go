package main

import "fmt"

type hunCode interface {
	printName()
}

type priscio struct {
	name string
}

func (p priscio) printName() {
	fmt.Println("Hello World")
}

func Print(hun hunCode) {
	fmt.Printf("O objeto por tras dessa interface hunCode é o objeto %T", hun)
	hun.printName()
}

func main() {
	obj := priscio{}

	Print(obj)

}
