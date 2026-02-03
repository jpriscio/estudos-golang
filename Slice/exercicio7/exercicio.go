package main

import "fmt"

func main() {

	pessoas := map[string][]string{
		"Joao_Priscio":  []string{"treinar", "dançar"},
		"Cesar_Augusto": []string{"trabalhar", "dançar"},
		"Ana_Alyce":     []string{"treinar", "dançar"},
	}

	for _, valor := range pessoas {
		fmt.Printf("- %s\n", valor)
	}

}

//crie um map com key tipo string e value tipo []string.
//Key deve conter nomes no formato sobrenome_nome
//Value deve conter os hobbies favoritos da pessoa
