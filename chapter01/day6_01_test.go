package chapter01

import (
	"encoding/json"
	"fmt"
	"testing"
)

//定义响应结构体
type Item struct {
	Title string
	URL   string
}

//定义响应结构体
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

//圆形结构体
type Circle struct {
	Radius float32
}

func (t *Circle) Area() float32 {
	return t.Radius * t.Radius * 3.14
}

func TestDay601Main(t *testing.T) {
	jsonStr := `{"data": {
	   "children" : [
	{ "data":{"title":"blog","url":"www.blog.com"}}] }}`
	res := Response{}
	//字符串转结构体
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res.Data.Children)

	r := Circle{10}
	fmt.Println(r.Area())
}
