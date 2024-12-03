package myio

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件并将内容转为字符串
func OpenFileToString(fastaFile string) string {
	// 打开文件
	file, err := os.Open(fastaFile)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		os.Exit(1)
	}
	defer file.Close()

	// 读取文件内容到字符串
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件内容出错:", err)
		os.Exit(1)
	}
	return content
}
