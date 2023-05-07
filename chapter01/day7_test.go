package chapter01

//interface
import (
	"fmt"
	"testing"
)

type PProgrammer struct {
	Name string
}

func (stu PProgrammer) Write() {
	fmt.Println("PProgrammer Write()")

}

// 技能接口
type SkillInterface interface {
	Write()
}

type Number int

func (x Number) Equal(i Number) bool {
	return x == i

}

func (x Number) LessThan(i Number) bool {
	return x < i

}

func (x Number) MoreThan(i Number) bool {
	return x > i

}

// 乘法
func (x *Number) Multiple(i Number) {
	*x = *x * i
}

// 除法
func (x *Number) Divide(i Number) {
	*x = *x / i
}

type NumberI interface {
	Equal(i Number) bool
	LessThan(i Number) bool
	MoreThan(i Number) bool
	Multiple(i Number)
	Divide(i Number)
}

func TestDay7Main(t *testing.T) {
	var pro PProgrammer //定义一个结构体 实现了 Write方法 就==实现的SkillInterface接口
	var a SkillInterface = pro
	a.Write()

	//类型赋值
	var x Number = 8
	var y = &x

	fmt.Println(y.Equal(8))

}
