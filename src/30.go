package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

func main() {
	event := Bootcamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	}

	fmt.Printf("Even on %s, location (%f, %f)",
		event.Date, event.Lat, event.Lon)
}
