package main

import (
	"fmt"
	"time"
)

type sensorInterface interface {
	ReadValue() int64
}

type sensor struct {
	path   string
	name   string
	values []*sensorValue
}

type sensorValue struct {
	timestamp time.Time
	value     int64
}

type sensorReader struct {
}

func (s *sensor) ReadValue() int64 {
	return int64(20400)
}

func (s *sensor) StoreReading(input sensorReading) {
	s.values = append(
		s.values[1:],
		&sensorValue{
			timestamp: time.Now(),
			value:     input.value.value,
		},
	)
	fmt.Println(s.values)
}

func NewSensor() *sensor {
	var values []*sensorValue
	for v := 0; v < 4; v++ {
		values = append(values, &sensorValue{value: int64(0)})
	}
	sensor := &sensor{
		path:   "yeah",
		name:   "somename",
		values: values,
	}
	return sensor
}

type sensorReading struct {
	name  string
	value sensorValue
}

func (s *sensor) Daemon(input chan sensorReading) {
	for {
		reading := <-input
		s.StoreReading(reading)
	}
	close(input)
}
