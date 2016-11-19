package main

import (
	"fmt"

	gm "github.com/merxer/mathlibrary" // download library use "go get -v <url>"
)

func main() {
	fmt.Println(gm.Add(1, 2))
	fmt.Println(gm.Mul(1, 2))
}
