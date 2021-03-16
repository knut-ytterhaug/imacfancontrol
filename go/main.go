package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://gobyexample.com/stateful-goroutines

func main() {

	f, err := os.Open("/sys/devices/platform/applesmc.768/temp31_input")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := make([]byte, 5)
	f.Read(b)
	d, _ := strconv.Atoi(string(b))

	if d > 40000 {
		fmt.Println("HOT AS hell:", d)
	}
}
