package main

import "fmt"

func main() {
	cities := []string{}
	cities = append(cities, "San Diego")
	cities = append(cities, "Bangkok")

	fmt.Println(cities)
}
