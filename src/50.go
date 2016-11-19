package main

import "fmt"

func main() {
	cities := make([]string, 2)
	cities[0] = "Tak"
	cities[1] = "Buriram"
	cities = append(cities, "San Diego", "Bangkok")

	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

	cities = append(cities, "Surin")

	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

}
