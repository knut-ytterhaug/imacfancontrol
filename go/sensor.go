package main

import "time"

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
	path  string
	value sensorValue
}

func (s *sensor) GetAverageValue() int64 {
	var total, count int64
	for _, v := range s.values {
		total += v.value
		count++
	}
	return int64(total / count)
}
