package chapter02

import (
	"fmt"
	"reflect"
	"testing"
)

// reflect 反射 能自描述和自控制
// reflect 反射包获取动态变量的值和类型ValueOf、TypeOf
func Test_day0202Main(t *testing.T) {
	var money float32 = 998.88
	pointer := reflect.ValueOf(&money)
	value := reflect.ValueOf(money)
	fmt.Println("value: ", reflect.ValueOf(&money))
	fmt.Println("type:  ", reflect.TypeOf(&money))

	convertPointer := pointer.Interface().(*float32)
	convertValue := value.Interface().(float32)

	fmt.Println(convertPointer, convertValue)
}
