package chapter01

//接口的组合
import (
	"fmt"
	"testing"
)

// 接口1
type Interface1 interface {
	Write(p []byte) (n int, err error)
}

// 接口2
type Interface2 interface {
	Close() error
}

// 接口Combine:组合
type InterfaceCombine interface {
	Interface1
	Interface2
}

type TestCombine struct {
	name []byte
}

func (this TestCombine) Write(p []byte) {
	fmt.Println("write1", p)
}

func (this TestCombine) Close() {
	fmt.Println("Close1")
}

func TestDay702Main(t *testing.T) {
	var tc TestCombine
	tc.Write([]byte("1"))
	tc.Close()
}
