package low

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

var (
	tables = []string{"user", "order"}
	count  = len(tables)
	ch     = make(chan bool, count)
)

//运行SQL语句
func SqlQuery(db *sql.DB, table string, ch chan bool) {
	fmt.Println("开始处理：", table)
	rows, _ := db.Query(fmt.Sprint("select * from % s", table))
	columus, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	//定义字段变量
	values := make([]sql.RawBytes, len(columus))
	//定义变量
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	totalValues := [][]string{}
	for rows.Next() {
		var s []string
		//把每行内容添加到scanArgs 也加values
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		for _, v := range values {
			s = append(s, string(v))
		}
		totalValues = append(totalValues, s)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	//导出到csv
	ExportToCSV(table+".csv", columus, totalValues)

}

func ExportToCSV(file string, columns []string, totalValues [][]string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	w := csv.NewWriter(f)
	for a, i := range totalValues {
		if a == 0 {
			w.Write(columns)
			w.Write(i)
		} else {
			fmt.Println(i)
		}
		w.Write(i)
	}
	w.Flush()
	fmt.Println("处理完毕", file)
}

func TestDay3Main(t *testing.T) {
	//db,err := sql.Open("mysql","root:123456@")
}
