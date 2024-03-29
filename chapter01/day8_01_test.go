package chapter01

import (
	"fmt"
	"testing"
)

var counter = func(n int) chan<- chan<- int {
	//定义一个chan<-int类型的chan
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				n++ //递增记数
			} else {
				request <- n //返回当前计数
			}
		}
	}()
	return requests
	//返回隐式转换到类型chan <-(chan <- int)

}(0)

func TestDay801Main(t *testing.T) {
	increase100 := func(done chan<- struct{}) {
		for i := 0; i < 100; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}
	done := make(chan struct{})
	go increase100(done)
	go increase100(done)
	<-done
	<-done

	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request)
}
