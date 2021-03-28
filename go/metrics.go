package main

import (
	"fmt"

	"github.com/VictoriaMetrics/metrics"
)

func AddMetrics() {

	a := metrics.NewCounter("supercounter{name=\"Yes\"}")
	a.Inc()
	fmt.Println(a)
}
