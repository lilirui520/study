package chapter01

import (
	"fmt"
	"testing"
)

func TestDay703Main(t *testing.T) {

	var a interface{} = func(i int) string {
		return fmt.Sprintf("d: %d ", i)
	}
	switch b := a.(type) {
	case nil:
		println("nil")
	case *int:
		println(*b)
	case func(int) string:
		println(b(88))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		println("unknown")
	}
}
