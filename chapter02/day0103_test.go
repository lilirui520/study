package chapter02

import (
	"fmt"
	"log"
	"os"
	"testing"
)

//函数使用注意事项
//recover()获取祖父级调用异常 直接调用无效

func Test_day0103Main(t *testing.T) {
	//正确使用
	// defer func() {
	// 	recover()
	// }()
	// panic(6)
	//闭包错误
	// for i := 0; i < 3; i++ {
	// 	defer func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	//循环内使用defer回导致资源延迟释放 解决办法就是用局部函数
	for i := 0; i < 3; i++ {
		func() {
			f, err := os.Open("./test.txt")
			if err != nil {
				log.Fatal(err)
			}
			var b []byte = make([]byte, 4096)
			//Go表示字符串的数据类型是string。我们知道字符串是由字节组成，而字节序列在Go是用byte 类型的切片（[]byte）表达，因此 string 类型和 []byte 是可以互相转化的。
			n, err := f.Read(b)
			if err != nil {
				fmt.Println("Open file Failed", err)
			}
			data := string(b[:n])
			fmt.Println(data)
			defer f.Close()
		}()
	}
}

//闭包错误的问题
