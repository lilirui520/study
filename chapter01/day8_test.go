package chapter01

import (
	"fmt"
	"testing"
	"time"
)

func Hello(c chan string) {
	name := <-c //从通道获取数据
	fmt.Println("hello", name)
}

// 一次性定时通知器
func TimeLong(d time.Duration) <-chan struct{} {
	//??不知道先这么写
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()
	return ch
}

func TestDay8Main(t *testing.T) {
	// 声明一个字符串类型的变量
	ch := make(chan string)
	go Hello(ch)
	ch <- "world"
	//关闭通道
	close(ch)

	fmt.Println("你好")
	<-TimeLong(time.Second)
	fmt.Println("过1s后继续显示")
	<-TimeLong(time.Second)
	fmt.Println("再过1s后显示")

	//time.Alter函数 也是可以使用超时机制实现中
}
