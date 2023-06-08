package chapter02

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDay0301(t *testing.T) {

	//待编码的字符串
	data := "Hello,Base64"
	data1 := "Hello"

	// Go语言支持标准和URL兼容的Base64

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	sEnc1 := base64.StdEncoding.EncodeToString([]byte(data1))

	fmt.Println(sEnc)
	fmt.Println(sEnc1)

	// Decode返回Base64字符串标识的字节
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	//使用与url兼容的base64 进行编码 解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.StdEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
