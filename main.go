package main

import (
	"springboot-hand-stand/myutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const databaseName  = "marvel"
func main() {
	str := "marvel_luobosi_xueliu_naitang"

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/marvel?charset=utf8")
	checkErr(err)

	println(myutil.ToUpperUnderlineWordInitialsForField(str))

	databaseNames, err := myutil.QueryDatabaseTableName(db, databaseName)

	for index, tableName := range databaseNames {

		println(index, tableName)
	}
}


func checkErr(err error)  {
	if err != nil {
		println(err)
	}
}
