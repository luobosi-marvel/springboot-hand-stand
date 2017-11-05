package myutil

import "strings"

/**
	将传进来的字符串的首字母转成大写
 */
func toUpperCaseFirstOne(str string) (s string) {
	return strings.ToUpper(str[0:1]) + str[1:]
}

/**
	将下划线后面的单词首字母变成大写，也就是驼峰形式
	这个方法使用与带下划线的字段值
 */
func ToUpperUnderlineWordInitialsForField(str string)(s string) {
	split := strings.Split(str, "_")
	s = split[0]
	for i := 1; i < len(split); i++ {
		split[i] = toUpperCaseFirstOne(split[i])
		s += split[i]
	}
	return
}

/**
	将下划线后面的单词首字母变成大写，也就是驼峰形式
	这个方法使用与带下划线的类名
 */
func ToUpperUnderlinedWordInitialsForClassName(str string)(s string) {
	split := strings.Split(str, "_")
	for i := 0; i < len(split); i++ {
		split[i] = toUpperCaseFirstOne(split[i])
		s += split[i]
	}
	return
}
