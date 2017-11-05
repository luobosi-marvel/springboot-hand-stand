package allstruct


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
	相当于 Java 中的 toString 方法
 */
func (d DataBaseStruct) String()  {
	/*	fmt.Println("Field: ", d.Field, "   Type: ", d.Type, "   Collation: ", d.Collation, "   Null: ", d.Null, "   Key: ", d.Key,
			"   Default: ", d.Default, "   Extra: ", d.Extra, "   Privileges: ", d.Privileges, "   Comment: ", d.Comment)*/
}