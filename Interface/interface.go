package main

import "fmt"

type circulo struct {
	R float64
}

type triangulo struct {
	base   float64
	altura float64
}

func (c circulo) area() float64 {
	return 3.14 * c.R * c.R
}

func (t triangulo) area() float64 {
	return t.base * t.altura
}

type areas interface {
	area() float64
}

func maiorarea(t, b areas) areas {
	if t.area() > b.area() {
		return t
	} else {
		return b
	}

}

func main() {

	bola := circulo{R: 1}

	tri := triangulo{base: 14, altura: 10}

	fmt.Println(maiorarea(bola, tri).area())

}
