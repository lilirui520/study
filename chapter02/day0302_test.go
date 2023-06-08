package chapter02

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

//http://127.0.0.1:8802/debug/pprof/

func TestDay0302(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe(":8802", nil))
	}()

	for {
		Sum("this is a test")
	}
}

func Sum(str string) string {
	date := []byte(str)
	sData := string(date)
	var sum = 0
	for i := 0; i < 10000; i++ {
		sum += i
		fmt.Println(sum)
	}
	return sData
}
