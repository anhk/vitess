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
	Parse("select trim(leading 'aab ' from s_name) as a from student where s_addr='世界';")
	Parse("select trim( 'aab ' from s_name) as a from student where s_addr='世界';")
	Parse("select trim( s_name) as a from student where s_addr='世界';")
	Parse("select trim(leading'aab ' from s_name) as a from student where s_addr='世界';")
}
