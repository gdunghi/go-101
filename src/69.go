package main

import "fmt"

func main() {
	input := 'b'

	switch input {
	case 'a', 'A':
		fmt.Println("Press A")
	case 'b', 'B':
		fmt.Println("Press B")
	default:
		fmt.Println("don't know")
	}
}
