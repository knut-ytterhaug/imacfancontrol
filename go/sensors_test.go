package main

import (
	"fmt"
	"testing"
)

type sensorMockReader struct {
}

func (s *sensorMockReader) ReadValue() int64 {
	return int64(1)
}

func TestReadValue(t *testing.T) {

	sensors := sensor{sensorInterface: &sensorMockReader{}}

	fmt.Println(sensors.ReadValue())

}
