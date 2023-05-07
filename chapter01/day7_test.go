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

func TestDay7Main(t *testing.T) {
	var pro PProgrammer //定义一个结构体 实现了 Write方法 就==实现的SkillInterface接口
	var a SkillInterface = pro
	a.Write()
}
