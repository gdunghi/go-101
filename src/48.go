package main

import "fmt"

func main() {
	cities := []string{}
	cities = append(cities, "San Diego", "Bangkok")

	fmt.Printf("%q\n", cities)
}
