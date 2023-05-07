package chapter01

//接口查询
import (
	"fmt"
	"testing"
)

func Len(array interface{}) int {
	var length int //数据的长度

	if array == nil {
		length = 0
	}

	//通过使用“接口类型.(type)”形式，加上switch ... case 语句来判断接口存储的类型

	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))

	case []float32:
		length = len(array.([]float32))

	default:
		length = 0
	}
	fmt.Println(length)

	return length
}

func TestDay701Main(t *testing.T) {
	Len([]int{1, 2, 3, 4})
	Len([]string{"1, 2, 3, 4", "ABC"})
}
