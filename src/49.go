package main

import "fmt"

func main() {
	cities := []string{}
	cities = append(cities, "San Diego", "Bangkok")

	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

	countries := make([]string, 42)
	fmt.Println(len(countries))

}
