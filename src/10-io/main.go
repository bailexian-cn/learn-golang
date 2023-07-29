package __io

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	// 获取当前目录
	//_, absPath, _, _ := runtime.Caller(0)
	//proPath := filepath.Dir(absPath)

	file := os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

	fileInfo, _ := file.Stat()

	fmt.Println(fileInfo.Name())

	fmt.Fprintln(file, "Hello World!")

	//// 读文件
	//path := proPath + "/file_store"
	//content := operations.ReadFile(path)
	//fmt.Println(content)
	//
	//// 写文件
	//writePath := proPath + "/writeTest"
	//operations.WriteFile(writePath, "xxxx")
	//fmt.Println(operations.ReadFile(writePath))

}
