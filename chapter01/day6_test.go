package chapter01

import (
	"fmt"
	"testing"
)

//结构体
type Programmer struct {
	Name   string
	GoodAt string
}

func TestDay6Main(t *testing.T) {
	var person Programmer
	person.Name = "jay"
	person.GoodAt = "JAVa"
	fmt.Println(person)

	person1 := Programmer{"jj", "php"}
	fmt.Println(person1)

	var person2 = new(Programmer)
	person2.Name = "1"
	person2.GoodAt = "python"
	var person3 = &Programmer{}
	(*person3).Name = "milo"
	(*person3).GoodAt = "GO"

	fmt.Println(person2)
}
