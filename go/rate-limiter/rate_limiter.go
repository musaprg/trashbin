package main

import (
	"math"
	"net/http"
	"time"
)

type RateLimiter struct {
	AllowRequestsPerSeconds int
	requestTimeQueue        []time.Time
}

func NewRateLimiter(allowRequestPerSeconds int) RateLimiter {
	return RateLimiter{allowRequestPerSeconds, []time.Time{}}
}

func (rt *RateLimiter) cleanQueue(ct time.Time) {
	for len(rt.requestTimeQueue) > 0 {
		v := rt.requestTimeQueue[0]

		d := int(math.Abs(v.Sub(ct).Seconds()))
		if d < 1 {
			break
		}

		rt.requestTimeQueue = rt.requestTimeQueue[1:]
	}
}

func (rt *RateLimiter) RateLimit(r http.Request) bool {
	ct := time.Now()

	rt.cleanQueue(ct)

	if len(rt.requestTimeQueue) > 0 {
		rate := (len(rt.requestTimeQueue) + 1)
		if rate >= rt.AllowRequestsPerSeconds {
			return false
		}
	}

	rt.requestTimeQueue = append(rt.requestTimeQueue, ct)

	return true
}
