package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	a, b := 5, 2
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
}
