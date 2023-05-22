package chapter02

//匿名函数
import (
	"fmt"
	"testing"
)

func MlutiFunc(a, b int) (func(int) int, func() int) {
	f1 := func(c int) int {
		return (a + b) * c
	}
	f2 := func() int {
		return 2 * (a + b)
	}
	return f1, f2
}

func Test_day0101Main(t *testing.T) {

	//不带参数的匿名函数
	f := func() {
		fmt.Println("不带参数的匿名函数")
	}
	f()
	fmt.Printf("%T  \n", f) //打印func()

	//带参数的匿名函数
	f1 := func(args string) {
		fmt.Println(args)
	}
	f1("带参数的匿名函数写法1")

	(func(args string) {
		fmt.Println(args)
	})("带参数的匿名涵数写法2")
	func(args string) {
		fmt.Println(args)
	}("带参数的匿名涵数写法2")

	//带返回值的匿名函数

	a := (func(args string) (re string) {
		return args + "+++"
	})
	fmt.Println(a("返回值的匿名涵数写法"))

	//返回多个匿名函数的函数
	f11, f22 := MlutiFunc(6, 8)
	fmt.Println(f11(3))
	fmt.Println(f22())

}
