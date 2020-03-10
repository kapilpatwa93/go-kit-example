package main

import (
	"github.com/go-kit/kit/log"
	kitpro "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

func main() {
	Here()
	var svc StringService
	svc = stringService{}
	logger := log.NewLogfmtLogger(os.Stderr)
	
	fieldKeys := []string {"method","error"}
	requestCount := kitpro.NewCounterFrom(prometheus.CounterOpts{
		Namespace:   "my_group",
		Subsystem:   "string_service",
		Name:        "request_count",
		Help:        "No of request received",
		ConstLabels: nil,
	}, fieldKeys)
	
	requestLatency := kitpro.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace:   "my_group",
		Subsystem:   "string_service",
		Name:        "request_latency_microseconds",
		Help:        "Total duration fo requests in microseconds",
		ConstLabels: nil,
		Objectives:  nil,
		MaxAge:      0,
		AgeBuckets:  0,
		BufCap:      0,
	}, fieldKeys)
	
	countResult := kitpro.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace:   "my_group",
		Subsystem:   "string_service",
		Name:        "count_result",
		Help:        "The result of each count methods",
		ConstLabels: nil,
		Objectives:  nil,
		MaxAge:      0,
		AgeBuckets:  0,
		BufCap:      0,
	}, fieldKeys)
	svc = instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           svc,
	}
	svc = loggingMiddleware{
		logger: logger,
		next:   svc,
	}



	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metric", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}


