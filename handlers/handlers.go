package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/krogertechnology/krogo/pkg/krogo"
)

func HandlerGauge(ctx *krogo.Context) (interface{}, error) {
	var labelValues []string

	path := ctx.Param("path")
	method := ctx.Param("method")

	labelValues = append(labelValues, path, method)

	check := ctx.Metrics.IncGauge("than_gauge_req", labelValues)
	if check != nil {
		return nil, errors.New("can not increase the counter of gauge")
	}

	return "successfully increment of gauge metric ", nil
}
func HandlerGauge1(ctx *krogo.Context) (interface{}, error) {
	var labelValues []string

	path := ctx.Param("path")
	method := ctx.Param("method")

	labelValues = append(labelValues, path, method)

	check := ctx.Metrics.DecGauge("than_gauge_req1", labelValues)
	if check != nil {
		return nil, errors.New("can not increase the counter of gauge")
	}

	return "successfully increment of gauge metric ", nil
}

func HandlerCounter(ctx *krogo.Context) (interface{}, error) {
	var labelValues []string

	path := ctx.Param("path")
	method := ctx.Param("method")

	labelValues = append(labelValues, path, method)

	check := ctx.Metrics.IncCounter("than_counter_req", labelValues)
	if check != nil {
		return nil, errors.New("can not increase the counter")
	}

	return "successfully increased the count of Counter metrics   path is " + path + "method is" + method, nil

}

func HandlerCounter1(ctx *krogo.Context) (interface{}, error) {
	var labelValues []string

	path := ctx.Param("path")
	method := ctx.Param("method")
	fmt.Println("path is")
	fmt.Println(path)
	fmt.Println("method is")
	fmt.Println(method)

	labelValues = append(labelValues, path, method)

	check := ctx.Metrics.IncCounter("than_counter_req1", labelValues)
	if check != nil {
		return nil, errors.New("can not increase the counter")
	}

	return "successfully increased the count of Counter metrics", nil
}

func HandlerHistogram(ctx *krogo.Context) (interface{}, error) {
	var labelValues []string

	path := ctx.Param("path")
	method := ctx.Param("method")
	value := ctx.Param("value")

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		// if key named value is not passed, set the value default
		v = 10
	}

	labelValues = append(labelValues, path, method)

	check := ctx.Metrics.Observe("than_hist_req", labelValues, v)
	if check != nil {
		return nil, errors.New("can not observe")
	}

	return "successfully observe of histogram metrics ", nil
}
