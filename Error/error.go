package main

import (
	"fmt"
)

type errorMensage struct {
	Message string
	Code    int
}

func (e *errorMensage) Error() string {
	return fmt.Sprintf("Erro foi: %s, codigo do erro: %v", e.Message, e.Code)
}

func teste() error {
	err := errorMensage{
		Message: "deu erro aqui",
		Code:    500,
	}
	return &err
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, teste()
	}
	return a / b, nil
}

func main() {

	resultado, err := divisao(50, 0)
	if err != nil {
		fmt.Printf("%#v", err)
	} else {
		fmt.Println(resultado)
	}
}
