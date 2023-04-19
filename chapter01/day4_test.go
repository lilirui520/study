package chapter01

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

func GetMap() {
	varMap := make(map[int64]string)
	varMap[1] = "mi"
	varMap[2] = "day day up"
	for _, num := range []int64{1, 2, 3, 4} {
		if _, ok := varMap[num]; ok {
			fmt.Printf("varMap 中包含key:%d \n", num)

		} else {
			fmt.Printf("varMap 不包含key:%d \n", num)
		}
	}
	for _, num := range []int64{1, 2, 3, 4} {
		//如果包含key，就知道value
		if v, s := varMap[num]; s {
			fmt.Printf("varMap 包含key :%d value: %s \n", num, v)
		} else {
			fmt.Printf("varMap 不包含key :%d value: %s \n", num, v)
		}
	}
}

func JsonTOmap() {
	jsonStr := `{
		"name": "sam",
		"goodAt": "Go Programming"
		}`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("json to map err:", err)
	}
	fmt.Println(mapResult)
}

func MapToJson() {
	instance := map[string]interface{}{
		"name":   "sam",
		"goodAt": "Go Programming",
	}
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		fmt.Println("map to json err:", err)
	}
	fmt.Println(string(jsonStr))
}

//map是无序的，需要排列切片
func MapSort() {
	var arr map[int]int
	arr = make(map[int]int, 5)
	arr[0] = 88
	arr[1] = 28
	arr[2] = 8
	arr[3] = 18

	var b []int
	fmt.Println("排序前的值如下：")
	for k, v := range arr {
		fmt.Println(k, v)
		b = append(b, v)
	}
	fmt.Println("排序后的值如下：")
	//按照数字排序
	sort.Ints(b)
	for k, v := range b {
		fmt.Println(k, v)
	}

}

func TestDay4Main(t *testing.T) {
	GetMap()
	JsonTOmap()
	MapToJson()
	MapSort()

}
