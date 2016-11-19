package main

import "fmt"

type Human struct {
	Age int
}

func newYear(h *Human) int {
	h.Age++
	return h.Age
}

func main() {
	me := &Human{Age: 20}
	fmt.Println(newYear(me))
	fmt.Println(me.Age)
}
