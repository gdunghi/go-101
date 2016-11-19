package main

import "fmt"

func main() {
	action := func() {
		fmt.Println("In Function")
	}

	action()
}
