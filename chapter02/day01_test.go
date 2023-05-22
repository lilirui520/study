package chapter02

//递归函数学习
import (
	"fmt"
	"testing"
)

// 判断是否是偶数
func Even(num int) bool {
	if num == 0 {
		return true
	}
	return Odd(RecursiveSign(num) - 1)
}

// 判断是否奇数
func Odd(num int) bool {
	if num == 0 {
		return false
	}
	return Even(RecursiveSign(num) - 1)

}

//递归签名

func RecursiveSign(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Test_dayMain(t *testing.T) {

	num7 := Even(7)
	n7 := Odd(7)
	fmt.Printf("7 是偶数？ %v\n", num7)
	fmt.Printf("7 是奇数？ %v\n", n7)
}
