package main

import (
	"fmt"
)

type User struct {
	Firstname, LastName string
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.Firstname, u.LastName)
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(u.Greeting())
}
