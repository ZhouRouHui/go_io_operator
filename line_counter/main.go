package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// os.Args 获取命令行的内容
	args := os.Args
	if len(args) < 2 {
		fmt.Println("请输入文件名称")
		return
	}
	filename := args[1]

	// 打开文件
	file, _ := os.Open(filename)
	defer file.Close()

	r := bufio.NewReader(file)

	var line int

	for {
		_, isPrefix, err := r.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}

		if !isPrefix {
			line++
		}
	}

	fmt.Printf("文件 '%s' 共 %d 行\n", filename, line)
}
