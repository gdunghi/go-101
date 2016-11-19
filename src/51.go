package main

import "fmt"

func main() {
	cities := make([]string, 2)
	cities = append(cities, "Tak", "Bangkok")

	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

	cities = append(cities, "Surin")

	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

}
