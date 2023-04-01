package low

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func GetUrl() {
	authKey := "895314XY"
	password := "24D6YB309ZCB"
	proxyServer := "219.151.125.106:31615"
	targetURL := "https://jshk.com.cn/ip"
	rawURL := fmt.Sprintf("http://%s:%s@%s", authKey, password, proxyServer)
	proxyUrl, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	req, _ := http.NewRequest("GET", targetURL, nil)
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Printf("request failed: %s\n", err)
		return
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(body))
	}
}

func TestDay4Main(t *testing.T) {
	fmt.Println("abc")
}
