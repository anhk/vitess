package main

import (
	"fmt"
	"github.com/anhk/vitess/vt/sqlparser"
)

func main() {
	str := "delete from student where a=3"
	st, err := sqlparser.Parse(str)
	if err != nil {
		panic(err)
	}

	fmt.Println(sqlparser.String(st))

}
