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
	这是规范
 */
func (d DataBaseStruct) String() string {
	return  "Field: " + d.Field + "   Type: " + d.Type + "   Comment: " + d.Comment

}
