package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() {
	fmt.Printf("%s say is %s\n", c.Name, c.Greeting)
}

func main() {
	Pessoa1 := Creature{
		Name:     "João",
		Greeting: "Bom dia",
	}

	Pessoa2 := Creature{
		Name:     "Maria",
		Greeting: "Boa noite",
	}

	Pessoa1.Greet()
	Pessoa2.Greet()
}
