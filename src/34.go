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
	event := &Bootcamp{}
	event.Name = "Go Hackathon"
	event.Lat = 34.012836
	event.Lon = -118.495338
	event.Date = time.Now()

	fmt.Printf("Event (%s) on %s, location (%f, %f)",
		event.Name, event.Date, event.Lat, event.Lon)
}
