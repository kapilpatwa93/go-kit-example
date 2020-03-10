package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercassssse",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	fmt.Println("inside logging uppercase")
	output, err = mw.next.Uppercase(s)
	return

}


func (mw loggingMiddleware) Count(s string) (output int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercasesss",
			"input", s,
			"n", output,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.Count(s)
	return

}

