package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// bufio 扩展包的示例

func main() {
	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)
	data, _ := bufReader.Peek(5) // bufReader.Peek() 只读不取

	// bufReader.Buffered() 返回缓冲区中当前总共有多少个字节的数据
	fmt.Println(data, bufReader.Buffered())

	str, _ := bufReader.ReadString(' ') // 从缓冲区中取出并读取，此时缓冲区中不再有已取出的数据
	fmt.Println(str, bufReader.Buffered())


	/**
	 * 写入
	 */
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "hello ")
	fmt.Fprint(w, "world\n")
	w.Flush() // 写进 buf 中的数据必须要在 flush 之后才会显示在设备上
}