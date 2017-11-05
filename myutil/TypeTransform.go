package myutil

import "strings"

/**
	将MySQL 中的数据类型转换成 Java 中的数据类型
 */
func TransformMysqlTypeToJavaType(mysqlType string) (javaType string) {
	if strings.EqualFold("int", mysqlType) {
		javaType = "Long"
	} else if strings.EqualFold("var", mysqlType) {
		javaType = "String"
	}  else if strings.EqualFold("cha", mysqlType) {
		javaType = "String"
	}  else if strings.EqualFold("tex", mysqlType) {
		javaType = "String"
	}  else if strings.EqualFold("tin", mysqlType) {
		javaType = "Integer"
	}  else if strings.EqualFold("sma", mysqlType) {
		javaType = "Integer"
	}  else if strings.EqualFold("med", mysqlType) {
		javaType = "Integer"
	}  else if strings.EqualFold("big", mysqlType) {
		javaType = "Long"
	}  else if strings.EqualFold("flo", mysqlType) {
		javaType = "Float"
	}  else if strings.EqualFold("dou", mysqlType) {
		javaType = "Double"
	}  else if strings.EqualFold("boo", mysqlType) {
		javaType = "Integer"
	}
	return
}