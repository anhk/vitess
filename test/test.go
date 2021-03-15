package main

import (
	"fmt"
	"github.com/anhk/vitess/vt/sqlparser"
)

func Parse(sql string) {

	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(sqlparser.String(stmt))
}

func main() {
	Parse("set session transaction isolation level read uncommitted;\n")
	Parse("start transaction")
	Parse("start transaction read only;")
	Parse("commit;")
	Parse("rollback;")
	Parse("savepoint x;")
	Parse("rollback to  x")
	Parse("release savepoint x;")
	Parse("set  autocommit  = 1;")
	Parse("select `s-idx` as `s-idx` from student ")
}
