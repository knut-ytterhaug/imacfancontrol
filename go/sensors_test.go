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

	values <- sensorReading{
		name: "foo",
		value: sensorValue{
			value:     int64(1),
			timestamp: time.Now()},
	}
	values <- sensorReading{
		name: "foo",
		value: sensorValue{
			value:     int64(1),
			timestamp: time.Now()},
	}
	values <- sensorReading{
		name: "foo",
		value: sensorValue{
			value:     int64(1),
			timestamp: time.Now()},
	}
	values <- sensorReading{
		name: "foo",
		value: sensorValue{
			value:     int64(1),
			timestamp: time.Now()},
	}
	time.Sleep(1 * time.Second)
	fmt.Println(sensor.values)
	fmt.Println("OMFG")
	//sensors.ReadValue = func(key string) int64 { return int64(1) }

}
