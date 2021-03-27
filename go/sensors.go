package main

import (
	"fmt"
)

type sensors struct {
	sensors *map[string]*sensor
}

func NewSensors() *sensors {
	sensors := &sensors{
		sensors: &map[string]*sensor{
			"a": {
				name: "horse",
			},
			"b": {
				name: "cow",
			},
		},
	}

	return sensors
}

func (s *sensors) Daemon(input chan sensorReading) {
	defer close(input)
	for {
		reading := <-input
		//		s.StoreReading(reading)
		fmt.Println(reading)
	}
}
