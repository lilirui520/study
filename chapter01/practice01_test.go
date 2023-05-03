package chapter01

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
)

type User struct {
	Uid      int
	Name     string
	Phone    string
	Email    string
	Password string
}

// 查询数据库并返回
func queryMultiRow() []User {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, dberr := sql.Open("mysql", dsn)
	if dberr != nil {
		panic(dberr)
	}

	rows, err := db.Query("select uid,name,phone,email from `user` where uid > ?", 0)
	if err != nil {
		fmt.Printf("query err: %v", err)
	}
	defer rows.Close()
	//循环读取结果集合的数据
	users := []User{}
	//rows.Next()遍历结果集
	for rows.Next() {
		u := new(User)
		//rows.Scan()将每一行的每一列保存到变量中。
		err := rows.Scan(&u.Phone, &u.Email, &u.Password)
		users = append(users, *u)
		if err != nil {
			fmt.Printf("scan failed,err %v \n", err)
			return nil
		}
	}
	return users
}

// 导出csv文件
func ExportCSV(filepath string, data [][]string) {
	fp, err := os.Create(filepath) // 创建文件句柄
	if err != nil {
		log.Fatalf("创建文件 %v", err)
	}
	defer fp.Close()
	fp.WriteString("\xEF \xBB \xBF")
	w := csv.NewWriter(fp)
	//输入变量
	w.WriteAll(data)
	w.Flush()
}

func prMain(t *testing.T) {
	//设置导入的文件名
	filename := "./exportUsers.csv"
	// 从数据库获取数据
	users := queryMultiRow()
	//定义二维数据标题
	column := [][]string{{"手机", "邮箱", "密码"}}

	for _, u := range users {
		str := []string{}
		str = append(str, u.Phone)
		str = append(str, u.Email)
		str = append(str, u.Password)
		column = append(column, str)
	}
	//导出
	ExportCSV(filename, column)
}
