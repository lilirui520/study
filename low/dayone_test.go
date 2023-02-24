package low

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func GetRandomString(l int, str []byte) string {
	//转换成数组
	bytes := str
	//结果
	result := []byte{}
	//随机：重点
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//循环多少次L=长度 重点
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}

func NewLenChars(length int, chars []byte) string {
	if length < 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length forNewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("ERROR")
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue //跳过数字以避免偏差
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}

		}
	}
}

func TestDayMain(t *testing.T) {
	str := "0123456789abcde"
	bytes := []byte(str)
	//伪随机
	b := GetRandomString(5, bytes)
	fmt.Println(b)
	//真随机

	a := NewLenChars(5, bytes)
	fmt.Println(a)
}
