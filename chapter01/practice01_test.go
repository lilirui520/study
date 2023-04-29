package chapter01

import (
	"fmt"
	"sort"
	"testing"
)

type User struct {
	Uid      int
	Name     string
	Phone    string
	Email    string
	Password string
}

func queryMultiRow() (User[]) {
	rows, err := db.Query("select uid,name,phone,email from `user` where uid > ?", 0)
	if err != nil {
		fmt.Printf("query err: %v", err)
	}
	defer rows.Close()
	for _, v := range rows {
		
	}
	append(User, )
	rows.

}

func prMain(t *testing.T) {
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
