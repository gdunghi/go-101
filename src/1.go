package main

import "fmt"

const (
	ConstExample = "Constant"
)

var (
	universityName = "CMRU"
	room           = 264
)

type Student struct {
	Firstname, Lastname string
	universityName      string
}

func (s *Student) print() {
	fmt.Println(s.Firstname, s.Lastname, s.universityName)
}

func main() {
	u := &Student{
		Firstname:      "THAWATCHAI",
		Lastname:       "SINGNGAM",
		universityName: universityName,
	}

	u.print()
}
