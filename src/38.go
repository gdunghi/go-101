package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

func (b *Bootcamp) show() {
	fmt.Printf("Event on %s, location (%f, %f)",
		b.Date, b.Lat, b.Lon)
}

func main() {

	event := &Bootcamp{
		34.012836,
		-118.495338,
		time.Now(),
	}

	event.show()

}
