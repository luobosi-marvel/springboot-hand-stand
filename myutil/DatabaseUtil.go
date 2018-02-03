package myutil

import (
	"database/sql"
	"springboot-hand-stand/allstruct"
	"fmt"
)
/**
	查询某个表的所有表结构
	param
	param tableName 		表名
	return tableStructure	表结构信息
 */
func QueryTableStructure(tableName string) (tableStructures []allstruct.DataBaseStruct, err error) {
	// 连接数据库
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/marvel?charset=utf8")
	defer db.Close()
	// 查询某个表的表结构的 sql
	sql := "SHOW FULL COLUMNS FROM " + tableName
	// 这里要注意 new 和 make 的区别，make 只能用来为 slice，map 或 chan 类型分配内存和初始化一个对象
	// new 的作用是初始化一个指向类型的指针
	tableStructure := new(allstruct.DataBaseStruct)
	tableStructures = make([]allstruct.DataBaseStruct, 10)

	stmt, err := db.Prepare(sql)
	res, err := stmt.Query()

	for res.Next() {
		res.Scan(&tableStructure.Field, &tableStructure.Type, &tableStructure.Collation, &tableStructure.Null,
				 &tableStructure.Key, &tableStructure.Default, &tableStructure.Extra, &tableStructure.Privileges, &tableStructure.Comment)
		tableStructures = append(tableStructures, *tableStructure)
		fmt.Println(tableStructure)
	}
	return
}

/**
	查询某个数据库中的所有表名

	@param db				操作数据的
	@param databaseName  	数据库的名字
	@return databaseNames[]  该切片包含了数据库里面所有的表名字
		   err				错误，如果 err 不为 nil 则说明该方法有错误
 */
func QueryDatabaseTableName(db *sql.DB, databaseName string) (databaseNames []string, err error) {
	defer db.Close()
	sql := "SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema = '" + databaseName + "'"
	stmt, err := db.Prepare(sql)
	stmt.Query()

	checkErr(err)
	row, err := db.Query(sql)
	checkErr(err)

	for row.Next() {
		var TableName string
		row.Scan(&TableName)
		databaseNames = append(databaseNames, TableName)
	}
	return
}

func checkErr(err error)  {
	if err != nil {
		println("出 bug 了", err)
	}
}
