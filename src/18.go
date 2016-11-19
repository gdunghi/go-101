package main

import "fmt"

func greater(a int, b int) bool {
	return a > b
}

func main() {
	a, b := 5, 2
	fmt.Printf("Value a(%d) > b(%d) ?\n", a, b)
	fmt.Println(greater(a, b))
}
