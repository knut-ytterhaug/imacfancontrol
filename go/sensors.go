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
	values struct {
		timestamp time.Time
		value     int64
	}
	sensorInterface
}

type sensorReader struct {
}

func (s *sensorReader) ReadValue() int64 {
	fmt.Println("READ THAT STUFF")
	return int64(20400)
}
