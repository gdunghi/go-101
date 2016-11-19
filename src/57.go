package main

import "fmt"

func main() {
	cities := map[string]int{
		"Tak":   1,
		"Surin": 2,
	}

	fmt.Printf("%#v\n", cities)

	fmt.Printf("%v\n", cities)

}
