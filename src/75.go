package main

import "fmt"

func runBeforeDIE() {
	println("Keep Log, Before exit program")
}

func main() {
	defer runBeforeDIE()
	fmt.Println("Hello, World")
}
