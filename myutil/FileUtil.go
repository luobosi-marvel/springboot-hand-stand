package myutil

import (
	"os"
	"fmt"
)

/**
	文件操作的工具方法
 */


/**
	根据文件路径读取文件，将文件存到字符串并返回出去
	param path 文件的全路径名包括文件名
 */
func RedaFileFunc(path string)(fileStr string, err error) {
	path = "C:\\Users\\dell1\\Desktop\\upload.html"
	fin,err := os.Open(path)
	defer fin.Close()
	if err != nil {
		fmt.Println(path,err)
		return
	}
	buf := make([]byte, 1024)
	for{
		// 读取文件，每次读取一个 buf
		n, _ := fin.Read(buf)
		// 读到0 则退出
		if 0 == n { break }
		// 将文件都读到字符串里面
		fileStr += string(buf[:n])
	}
	return
}


/**
	将字符串写到指定目录的中去
	param path		 文件要生成的路径（不包括文件名）
	param fileStr	 要写入文件的字符串的内容
	param fileName	 要生成的文件的名字
 */
func WriteFileFunc(path string, fileStr string, fileName string)(err error) {

	path = "C:\\Users\\dell1\\Desktop\\" + fileName + ".java"
	// 如果该文件不存在则系统会创建该文件
	file_out,err := os.Create(path)
	defer file_out.Close()
	if err != nil {
		fmt.Println(path,err)
		return
	}
	// 写入字符串
	file_out.WriteString(fileStr)
	return
}

/**
	判断文件是否存在
 */
func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

