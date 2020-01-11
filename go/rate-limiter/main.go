package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	nr := NewRateLimiter(10)
	for i := 0; i < 1000; i++ {
		time.Sleep(20 * time.Millisecond)
		fmt.Println(nr.RateLimit(http.Request{}))
	}
}
