package main

import "fmt"

func main() {
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)

	fmt.Println(mySlice[1:4]) // value is [lo:hi-1]

	fmt.Println(mySlice[:4]) // value is [0:hi-1]

	fmt.Println(mySlice[4:]) // value is [4:len(s)]
}
