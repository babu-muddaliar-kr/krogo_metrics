package main

import (
	//"github.com/krogertechnology/krogo/examples/using-metrics/handlers"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"learning/handlers"
)

func main() {
	// create the application object
	k := krogo.New()
	k.Server.HTTP.Port = 9030
	k.Server.ValidateHeaders = false
	k.Server.MetricsPort = 2221

	labelName := []string{"path", "method"}

	_ = k.NewCounter("than_counter_req", "test counter", labelName)
	k.GET("/countermetrics", handlers.HandlerCounter)

	_ = k.NewCounter("than_counter_req1", "test counter1", labelName)
	k.GET("/countermetrics1", handlers.HandlerCounter1)

	_ = k.NewGauge("than_gauge_req", "test Gauge", labelName)
	k.GET("/gaugemetrics", handlers.HandlerGauge)

	_ = k.NewGauge("than_gauge_req1", "test Gauge1", labelName)
	k.GET("/gaugemetrics1", handlers.HandlerGauge1)

	//
	_ = k.NewHistogram("than_hist_req", "test Histogram", labelName,
		[]float64{.001, .003, .005, .01, .025, .05, .1, .2, .3, .4, .5, .75, 1, 2, 3, 5, 10, 30})
	k.GET("/hello-histogram", handlers.HandlerHistogram)

	// start the server

	k.Start()

}

