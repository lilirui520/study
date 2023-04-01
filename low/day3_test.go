package low

import (
	"database/sql"
	"fmt"
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

}

func TestDay3Main(t *testing.T) {

}
