package main

import (
	"fmt"
	"testing"
	"time"
)

//go:generate go-mockgen -f github.com/example/package -i ExampleInterface -o mock_example_interface_test.go

type sensorMockReader struct {
	ReadValue func(string) int64
}

func TestReadValue(t *testing.T) {

	sensor := NewSensor()
	values := make(chan sensorReading, 3)

	go sensor.Daemon(values)

	for i := 0; i < 20; i++ {
		values <- sensorReading{
			value: sensorValue{
				value:     int64(10),
				timestamp: time.Now(),
			},
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Println(sensor.GetAverageValue())

}
