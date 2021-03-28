package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type sensors struct {
	sensors *map[string]*sensor
}

func NewSensors() *sensors {
	var files []string

	root := "/sys/devices/platform/applesmc.768/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

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
