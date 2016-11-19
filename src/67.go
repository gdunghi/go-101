package main

import "fmt"

func main() {
	point := 75

	switch point {
	case 50:
		fmt.Println("not pass")
	case 75:
		fmt.Println("pass")
	default:
		fmt.Println("Excellent")
	}
}
