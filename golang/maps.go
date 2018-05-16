package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"bob": 5}

	m["marley"] = 22

	for key, value := range m {
		fmt.Println(key, value)
	}

}
