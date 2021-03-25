package main

import (
	"fmt"
)

// https://gobyexample.com/stateful-goroutines

func main() {

	sensor := NewSensor()

	fmt.Println(sensor)
}
