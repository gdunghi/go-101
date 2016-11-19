package main

import "fmt"

func runBeforeDIE() {
	println("Keep Log, Before exit program")
}

func runPanic() {
	panic("Crash.....")
}

func main() {
	defer runBeforeDIE()
	runPanic()
	fmt.Println("Hello, World")
}
