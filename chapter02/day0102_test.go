package chapter02

//匿名函数
import (
	"fmt"
	"testing"
)

func Func() func(string) string {
	a := "+++"
	return func(args string) string {
		a += args
		return a
	}
}

func Funcc() (r int) {
	defer func() {
		r += 3
	}()
	return 0
}

func Funccc() []func() {
	b := make([]func(), 2, 2)
	for i := 0; i < 2; i++ {
		b[i] = (func(j int) func() {
			return func() {
				fmt.Println(&j, j)
			}
		})(i)
	}
	return b
}

func Test_day0102Main(t *testing.T) {

	a := Func()
	b := a("你好！")
	c := a("世界")
	fmt.Println(b)
	fmt.Println(c)

	//defer后进先出
	fmt.Println(Funcc())

	//定一个闭包
	arr := Funccc()
	arr[0]()
	arr[1]()
}
