package springboottemplate

import (
	"springboot-hand-stand/myutil"
	"springboot-hand-stand/allstruct"
)

/**
	创建 DO 类
	param tableName				数据库中表名（也就是要生成实体的名字）
	param projectPackage		项目主包名（生成包名的时候需要使用）
	param tableStructures		表结构信息
	param path					DO 类生成的文件的目录
 */
func CreateDO(tableName string, projectPackage string, tableStructures []allstruct.DataBaseStruct, path string)  {

	// 用于存放创建的整个类的字符串
	var templateStr string

	templateStr += projectPackage + "\r\n"
	templateStr += "\r\n"
	// 这里添加要导入的包文件
	templateStr += projectPackage + ".domain.base.Base;"
	templateStr += "lombok.Data;"
	templateStr += "\r\n"
	templateStr += "/**\r\n"
	// 这里就是用来替换类名
	templateStr += " * " + myutil.ToUpperUnderlinedWordInitialsForClassName(tableName) + "DO 实体类\r\n"
	templateStr += " * \r\n"
	templateStr += " * @author 萝卜丝\r\n"
	templateStr += " */\r\n"
	// 这里加上 @Data 注解，就可以少写一些 Getter 和 Setter 方法
	templateStr += "@Data\r\n"
	// 这里是类名,也需要替换
	templateStr += "public class " + myutil.ToUpperUnderlinedWordInitialsForClassName(tableName) + "DO extends Base{\r\n"
	templateStr += "\r\n"
	//////////////////////////////////////////////////////////////////// 这里是需要根据数据库里面的结果生成字段

	// 获取表里面的所有的数据 field，type，comment
	for _, tableStructure := range tableStructures{
		// 获取字段名，并将数据库中的字段名转换成驼峰形式
		field := myutil.ToUpperUnderlineWordInitialsForField(tableStructure.Field)
		// 截取Type 的前三个字符然后再转换成 Java 类型
		category := myutil.TransformMysqlTypeToJavaType(tableStructure.Type[0:3])
		// 获取 comment
		comment := tableStructure.Comment

		templateStr += "	/**\r\n"
		templateStr += "	 * " + field + " " + comment +"\r\n"
		templateStr += "	 */\r\n"
		templateStr += "	private " + category + " " + field + ";\r\n"
		templateStr += "\r\n"
	}
	templateStr += "}\r\n"

	// 将该字符串写入指定的文件目录下面
}