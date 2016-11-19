package main

import "fmt"

const Pi = 3.14
const (
	StatusOK      = 200
	StatusCreated = 201
)

func main() {
	const Greeting = "สวัสดีครับ"

	fmt.Println(Pi)
	fmt.Println(StatusOK)
	fmt.Println(StatusCreated)
	fmt.Println(Greeting)
}
