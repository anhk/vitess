package main

import (
	"fmt"
	"github.com/anhk/vitess/vt/sqlparser"
)

func Parse(str string) {
	fmt.Println("==========================")
	fmt.Println(str)
	st, err := sqlparser.Parse(str)
	if err != nil {
		panic(err)
	}

	fmt.Println(sqlparser.String(st))
}

func main() {
	//CONVERT openb expression USING charset closeb
	Parse("select s_name from student where s_phone = any(select s_phone from student);")
	Parse("select s_name from student where s_phone = all(select s_phone from student);")
	Parse("select s_name from student where s_phone = some(select s_phone from student);")
}
