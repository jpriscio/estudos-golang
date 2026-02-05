package main

import (
	"fmt"
)

type errorMensage struct {
	Mensage string
	Code    int
}

func (e *errorMensage) Error() string {
	return fmt.Sprintf("Erro foi: %s, codigo do erro: %v", e.Mensage, e.Code)
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &errorMensage{Mensage: "divisor igual a zero", Code: 505}
	}
	return a / b, nil
}

func main() {

	resultado, err := divisao(50, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
