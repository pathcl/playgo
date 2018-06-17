package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("slices.go")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer f.Close()
	b := make([]byte, 1024)
	n, _ := f.Read(b)

	fmt.Printf("%d: % s\n", n, b)
}
