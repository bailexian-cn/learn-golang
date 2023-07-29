package __io

import (
	"fmt"
	"os"
)

func WriteFile(absPath string, content string) {
	file, err := os.Create(absPath)
	if err != nil {
		fmt.Println("create file error: ", err)
		return
	}

	defer file.Close()

	file.WriteString(content)
}
