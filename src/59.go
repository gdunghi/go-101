package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	for key, value := range m {
		fmt.Printf("key(%s), value(%v)", key, value)
	}
}
