package main

// io 扩展包的一些操作示例

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, n int) ([]byte, error) {
	p := make([]byte, n)
	n, err := reader.Read(p)
	if err != nil {
		return p, err
	}
	return p[:n], nil
}

// 读取一个字符串
func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)

	fmt.Println(data)
}

// 读取命令行输入的内容
func sampleReadFromStdin() {
	fmt.Println("Please input from stdin:")
	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

// 从文件中读取
func sampleReadFromFile() {
	file, _ := os.Open("basic_io/main.go")
	defer file.Close()

	data, _ := ReadFrom(file, 10)
	fmt.Println(string(data))
}

func main() {
	sampleReadFromString()
	sampleReadFromStdin()
	sampleReadFromFile()
}