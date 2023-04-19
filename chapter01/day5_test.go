package chapter01

import (
	"fmt"
	"sort"
	"testing"
)

func TestDay5Main(t *testing.T) {
	m := make(map[int]string)
	m[0] = "I"
	m[2] = "GO"
	m[1] = "LOVE"
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("key:", k, "value:", m[k])
	}
}
