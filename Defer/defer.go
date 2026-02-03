package main

import "fmt"

func main() {

	defer fmt.Println("ultimo")
	defer fmt.Println("meio")
	defer fmt.Println("primeiro")

}
