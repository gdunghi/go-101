package main

import "fmt"

func main() {
	cities := map[string]int{
		"Tak":   1,
		"Surin": 2,
	}

	for key, value := range cities {
		fmt.Printf("%s is number %d\n", key, value)
	}
}
