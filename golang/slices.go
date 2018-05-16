package main

import (
	"fmt"
)

func printer(w []string) {
	for _, word := range w {
		fmt.Println(word)
	}
}

func main() {
	words := []string{"the", "quick", "brown", "fox"}
	fmt.Println(len(words), cap(words))
	words = append(words, "hola")
	fmt.Println(len(words), cap(words))
	printer(words)

	//copy slices
}
