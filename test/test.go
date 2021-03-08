package main

import (
	"fmt"
	"github.com/anhk/vitess/vt/sqlparser"
)

func main() {
	stmt, err := sqlparser.Parse("select s_phone /** 33 **/ from student where s_id = 1")
	if err != nil {
		panic(err)
	}

	stmtCloned := stmt.CloneAsStatement()
	stmtCloned.(*sqlparser.Select).Where = nil

	fmt.Println(sqlparser.String(stmtCloned))
	fmt.Println(sqlparser.String(stmt))

}
