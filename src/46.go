package main

import "fmt"

func main() {
	mySlice := make([]int, 3)
	mySlice[0] = 2
	mySlice[1] = 3
	mySlice[2] = 5

	fmt.Println(mySlice)
}
