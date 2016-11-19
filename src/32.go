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
	event := new(Bootcamp)
	event.Lat = 34.012836
	event.Lon = -118.495338
	event.Date = time.Now()

	fmt.Printf("Even on %s, location (%f, %f)",
		event.Date, event.Lat, event.Lon)
}
