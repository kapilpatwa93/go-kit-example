package main

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           StringService
}

func (mw instrumentingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	fmt.Println("inside metrics ")
	output, err = mw.next.Uppercase(s)
	return
}

func (mw instrumentingMiddleware) Count(s string) (count int){
	defer func(begin time.Time) {
		lsv := [] string{"method", "count", "error", "false"}
		mw.requestCount.With(lsv...).Add(1)
		mw.requestLatency.With(lsv...).Observe(time.Since(begin).Seconds())

	}(time.Now())
	count = mw.next.Count(s)
	return
}
