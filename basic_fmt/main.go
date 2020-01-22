package main

import (
	"fmt"
	"os"
)

// 内置 format 包学习
func main() {
	fmt.Println("hello world", 2, 3.3, true) // 自带换行，可同时打印多个值

	fmt.Printf("my name is %s\n", "loedan") // 格式化输出，不自带换行

	myAge := fmt.Sprintf("my age is %d\n", 30) // 将格式化后的内容赋值给一个变量
	fmt.Print(myAge)

	// 所有开头的 F 代表文件的意思，所有 print 后面的 f 代表 format 格式化的意思
	// 在类 unix 系统中所有的东西都是 File 文件，所以在这里以 F 代表文件
	fmt.Fprint(os.Stdout, "kevin\n")

	fmt.Printf("%v\n", Data{})
}

type Data struct {

}

// 定义这个方法会使这个结构体实现 Stringer 接口，String() 方法相当于 php 和 java 的面向对象中的 toString() 方法
func (self Data) String() string {
	return "data"
}