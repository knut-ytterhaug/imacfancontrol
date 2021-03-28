package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

func main() {

	sensors := NewSensors()

	fmt.Println(sensors)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		metrics.WritePrometheus(w, true)
	})

	AddMetrics()
	http.ListenAndServe(":8080", nil)

	time.Sleep(10 * time.Second)
}
