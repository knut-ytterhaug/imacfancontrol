package main

import (
	"testing"
	"time"
)

//go:generate go-mockgen -f github.com/example/package -i ExampleInterface -o mock_example_interface_test.go

type sensorMockReader struct {
	ReadValue func(string) int64
}

/*
func TestReadValue(t *testing.T) {

	sensor := NewSensor()
	values := make(chan sensorReading, 1)

	go sensor.Daemon(values)

	for i := 50000; i < 60000; i += 1000 {
		values <- sensorReading{
			value: sensorValue{
				value:     int64(i),
				timestamp: time.Now(),
			},
		}
	}

	time.Sleep(1 * time.Second)
	if r := sensor.GetAverageValue(); r != 57500 {
		t.Errorf("Average failed. Expected %v, got %v", 57500, r)
	}

}*/

func TestGetAverageValue(t *testing.T) {

	sensor := NewSensor()

	for i := 50000; i < 60000; i += 1000 {
		sensor.StoreReading(sensorReading{
			value: sensorValue{
				value:     int64(i),
				timestamp: time.Now(),
			},
		})
	}

	if r := sensor.GetAverageValue(); r != 57500 {
		t.Errorf("Average failed. Expected %v, got %v", 57500, r)
	}

}
