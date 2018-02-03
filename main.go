package main

import (
	"springboot-hand-stand/myutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const databaseName  = "marvel"

func main() {
	//
	str := "marvel_luobosi_xueliu_naitang"

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/marvel?charset=utf8")
	checkErr(err)

	println(myutil.ToUpperUnderlineWordInitialsForField(str))
	// 查询某个数据库中的所有表名
	databaseNames, err := myutil.QueryDatabaseTableName(db, databaseName)

	for index, tableName := range databaseNames {
		println(index, tableName)
	}
	// 查询表结构信息
	tableStructures, err := myutil.QueryTableStructure("account")

	if err != nil {
		println("出 bug 了")
	}
	println("长度为: ", len(tableStructures))

}

func checkErr(err error)  {
	if err != nil {
		println(err)
	}
}
