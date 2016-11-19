package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Name string
	}
	Bootcamp struct {
		Event
		Lat, Lon float64
		Date     time.Time
	}
)

func main() {

	event := &Bootcamp{
		Event{"Go Hackathorn"},
		34.012836,
		-118.495338,
		time.Now(),
	}

	fmt.Printf("Event (%s) on %s, location (%f, %f)",
		event.Name, event.Date, event.Lat, event.Lon)
}
