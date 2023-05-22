package chapter01

import (
	"fmt"
	"testing"
	"time"
)

//1min内处理的请求数不会超过100

type Request interface{}

func handle(r Request) {
	fmt.Println(r.(int))
}

//限速时间1min
const RateLimitPeriod = time.Minute

//100限速
const RateLimit = 100

//处理多个请求的函数
func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func TestDay802Main(t *testing.T) {
	requests := make(chan Request)
	go handleRequests(requests)
	for i := 0; ; i++ {
		requests <- i
	}

}
