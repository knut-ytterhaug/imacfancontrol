package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type sensors struct {
	sensors *map[string]*sensor
}

func NewSensors() *sensors {
	var files []string

	root := "/sys/devices/platform/applesmc.768/"
	i := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "_label") {
			files = append(files, path)
			i = i + 1
			fmt.Println("how many times am I here?", i, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	print(files)

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
