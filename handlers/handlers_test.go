package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"github.com/krogertechnology/krogo/pkg/krogo/request"
)

func initializeHandlersTest(t *testing.T) *krogo.Krogo {
	k := krogo.New()
	k.NewCounter("than_counter_req", "test-help", []string{"path", "method"})
	k.NewGauge("than_gauge_req", "test-help", []string{"path", "method"})
	k.NewHistogram("than_hist_req", "test-help", []string{"path", "method"}, []float64{.001, .003, .005, .01, .025, .05, .1, .2, .3, .4, .5, .75, 1, 2, 3, 5, 10, 30})
	return k
}

func TestHandler(t *testing.T) {
	k := initializeHandlersTest(t)

	path := "hello-gauge"
	method := "get"

	req := httptest.NewRequest("GET", "http://dummy", nil)
	q := req.URL.Query()
	q.Add("method", method)
	q.Add("path", path)
	req.URL.RawQuery = q.Encode()

	r := request.NewHTTPRequest(req)
	context := krogo.NewContext(nil, r, k)

	// for counter handler
	_, err := HandlerCounter(context)
	assert.Equal(t, false, err != nil)

	// for gaugeHandler
	_, err = HandlerGauge(context)
	assert.Equal(t, false, err != nil)

	// for histogramHandler
	_, err = HandlerHistogram(context)
	assert.Equal(t, false, err != nil)
}

func initializeErrorHandlersTest(t *testing.T) *krogo.Krogo {
	k := krogo.New()

	k.NewCounter("thanTest_counter_request", "test-help ", []string{"path", "method"})
	k.NewGauge("thanTest_gauge_request", "test-help ", []string{"path", "method"})

	k.NewHistogram("thanTest_histogram_request", "test-help", []string{"path", "method"}, []float64{.001, .003, .005, .01, .025, .05, .1, .2, .3, .4, .5, .75, 1, 2, 3, 5, 10, 30})
	return k
}

func TestHandlerError(t *testing.T) {
	k := initializeErrorHandlersTest(t)

	req := httptest.NewRequest("GET", "http://dummyerror", nil)
	q := req.URL.Query()
	path := "hello-gauge"
	method := "get"
	q.Add("method", method)
	q.Add("path", path)
	req.URL.RawQuery = q.Encode()

	r := request.NewHTTPRequest(req)
	context := krogo.NewContext(nil, r, k)

	// for counter handler, error will come if it is not registered. Here name is used which is not registered
	_, err := HandlerCounter(context)
	assert.Equal(t, errors.New("can not increase the counter"), err)

	// for gauge  handler, error will come if it is not registered. Here name is used which is not registered
	_, err = HandlerGauge(context)
	assert.Equal(t, errors.New("can not increase the counter of gauge"), err)

	_, err = HandlerHistogram(context)
	assert.Equal(t, errors.New("can not observe"), err)
}
