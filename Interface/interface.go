package main

import "fmt"

type circulo struct {
	pi float64
	R  float64
}

type triangulo struct {
	base   float64
	altura float64
}

func (c circulo) area() float64 {
	area := 0.0
	c.pi = 3.14
	area = c.pi * (c.R * 2)
	return area
}

func (t triangulo) area() float64 {
	area := t.base * t.altura / 2
	return area
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

	bola := circulo{
		pi: 3.14,
		R:  10,
	}

	tri := triangulo{
		base:   14,
		altura: 10,
	}

	fmt.Println(maiorarea(bola, tri).area())

}
