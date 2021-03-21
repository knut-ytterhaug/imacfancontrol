package main

import (
	"fmt"
	"testing"
)

//go:generate go-mockgen -f github.com/example/package -i ExampleInterface -o mock_example_interface_test.go

type sensorMockReader struct {
	ReadValue func(string) int64
}

func TestReadValue(t *testing.T) {

	sensors := sensor{}

	values := make(chan string, 3)
	values <- "foo"
	values <- "bar"
	values <- "baz"

	fmt.Println(sensors.ReadValue())
	sensors.ReadValue = func(key string) int64 { return int64(1) }

}
