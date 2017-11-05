package util

import (
	"database/sql"
	"springboot-hand-stand/allstruct"
)
/**
	查询某个表的所有表结构
	param
	param tableName 		表名
	return tableStructure	表结构信息
 */
func QueryTableStructure(db *sql.DB, tableName string) (tableStructure map[string]interface{}, err error) {
	defer db.Close()
	sql := "SHOW FULL COLUMNS FROM " + tableName
	stmt, err := db.Prepare(sql)
	stmt.Query()

	checkErr(err)
	row, err := db.Query(sql)
	checkErr(err)

	for row.Next() {
		var Field string
		var Type string
		var Collation interface{}
		var Null interface{}
		var Key interface{}
		var Default interface{}
		var Extra interface{}
		var Privileges string
		var Comment string
		row.Scan(&Field, &Type, &Collation, &Null, &Key, &Default, &Extra, &Privileges, &Comment)

		tableStructure["Field"] = Field
		tableStructure["Type"] = Type
		tableStructure["Collation"] = Collation
		tableStructure["Null"] = Null
		tableStructure["Key"] = Key
		tableStructure["Default"] = Default
		tableStructure["Extra"] = Extra
		tableStructure["Privileges"] = Privileges
		tableStructure["Comment"] = Comment
	}
	return
}

/**
	查询某个数据库中的所有表名
	param db				操作数据的
	param databaseName  	数据库的名字
	return databaseNames[]  该切片包含了数据库里面所有的表名字
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
		println(err)
	}
}
