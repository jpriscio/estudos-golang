package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	cont := 0

	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {

			if c1[i] != c2[j] {
				cont++
			}
		}
	}

	fmt.Println(cont)

}
