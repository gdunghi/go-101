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
	eventName := &Event{Name: "Go Hackathorn"}

	event := &Bootcamp{
		Event: *eventName,
		Lat:   34.012836,
		Lon:   -118.495338,
		Date:  time.Now(),
	}

	fmt.Printf("Event (%s) on %s, location (%f, %f)",
		event.Name, event.Date, event.Lat, event.Lon)
}
