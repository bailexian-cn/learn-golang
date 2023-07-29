package __io

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(absolutePath string) string {
	file, err := os.Open(absolutePath)

	if err != nil {
		fmt.Println("open file err: ", err)
		return err.Error()
	}
	// 读取文件缓冲区
	var buf [128]byte
	// 文件内容字节数组
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 文件读取完毕
			break
		}

		if err != nil {
			fmt.Println("read file err: ", err)
			return err.Error()
		}

		content = append(content, buf[:n]...)
	}

	return string(content)
}
