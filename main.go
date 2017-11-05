package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
	用来存储 MySQL 的表的数据结构
 */
type DataBaseStruct struct {
	Field string
	Type string
	Collation interface{}
	Null interface{}
	Key interface{}
	Default interface{}
	Extra interface{}
	Privileges string
	Comment string
}

/**
	相当于 Java 中的 头String
 */
func (d DataBaseStruct) String()  {
/*	fmt.Println("Field: ", d.Field, "   Type: ", d.Type, "   Collation: ", d.Collation, "   Null: ", d.Null, "   Key: ", d.Key,
		"   Default: ", d.Default, "   Extra: ", d.Extra, "   Privileges: ", d.Privileges, "   Comment: ", d.Comment)*/
}
func main() {

	db, err := sql.Open("mysql", "root:root@tcp(47.94.172.234)/fire_ads?charset=utf8")
	checkErr(err)
	sql := "SHOW FULL COLUMNS FROM ad"
	fmt.Println(db, sql)
	queryTableStructure(db, sql)

}
/**
	查询某个表的所有表结构
 */
func queryTableStructure(db *sql.DB, sql string)  {
	defer db.Close()
	stmt, err := db.Prepare(sql)
	stmt.Query()
	checkErr(err)
	row, err := db.Query(sql)
	checkErr(err)
	// column, _ := row.Columns()

	for row.Next() {
		dataBaseStruct := new(DataBaseStruct)
		row.Scan(&dataBaseStruct.Field, &dataBaseStruct.Type, &dataBaseStruct.Collation, &dataBaseStruct.Null, &dataBaseStruct.Key,
		  		 &dataBaseStruct.Default, &dataBaseStruct.Extra, &dataBaseStruct.Privileges, &dataBaseStruct.Comment)
		fmt.Println(dataBaseStruct)
	}
}

func checkErr(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}
