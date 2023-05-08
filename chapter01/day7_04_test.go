package chapter01

//接口的多态
import (
	"fmt"
	"testing"
)

// Message 是定义通知类行为的接口
type Message interface {
	sending()
}

// 定义programmer 及 sending方法

type PPprogrammer struct {
	name  string
	phone string
}

func (u *PPprogrammer) sending() {
	fmt.Printf("sending programmer %s,%s \n", u.name, u.phone)
}

// 定义student及student.message方法
type Student struct {
	name  string
	phone string
}

func (s *Student) sending() {
	fmt.Printf("sending Student %s,%s \n", s.name, s.phone)
}

// sendMessage方法接收一个实现了message接口的值，并发通知
func sendMessage(n Message) {
	n.sending()
}

func TestDay704Main(t *testing.T) {
	//创建一个programmer值并传给sendmessage

	bill := PPprogrammer{"jack", "123456"}

	sendMessage(&bill)

	lee := Student{"lee", "789123"}

	sendMessage(&lee)
}
