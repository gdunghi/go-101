package main

import "fmt"

var (
	goIsFun bool   = true
	name    string = "Pat"
)

func main() {
	const format = "%T(%v)\n"
	fmt.Printf(format, goIsFun, goIsFun)
	fmt.Printf(format, name, name)
}
