package main

import "fmt"

func main() {
	//常亮用const标识，在程序编译器确定值，运行时无法改变值。
	//const constantName = value
	const PI = 3.141592653
	const i = 100
	const prefix = "za_"

	fmt.Printf("pi:%f,i:%d,prefix:%s", PI, i, prefix)
}
