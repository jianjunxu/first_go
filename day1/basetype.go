package main

import (
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	//布尔值 bool 默认false
	//整型 int,uint
	//复数 complex64,complex128 var c complex64 = 5+5i
	//浮点型 float32 float64
	//字符串
	//错误类型 error
	//array var arr [n]type
	//slice
	//map map[keyType]valueType

	//make new 操作
	//make用于内建类型（map,slice和channel）的内存分配，new用于各种类型的内存分配

	err := errors.New("error sksalj")
	if err != nil {
		fmt.Println(err)
	}

	//分组定义
	var (
		i = 100
		pi = 3.14
		prefix = "hello_"
	)
	fmt.Printf("i:%d,pi:%f,prefix:%s\r\n", i, pi, prefix)

	const (
		x = iota
		y = iota
		z = iota
		w
	)
	fmt.Printf("x:%d,y:%d,z:%d,w:%d\r\n", x, y, z, w)

	arr := [...]int{1, 2, 3}
	fmt.Printf("arr[1]:%d\r\n", arr[1])

	doubleArr := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("doubleArr[1][2]:%d\r\n", doubleArr[1][2])

	slice := []byte{'a', 'b', 'c', 'd'}
	nArr := slice[2:]
	fmt.Printf("slice[1]:%d,nArr[1]:%d\r\n", slice[1], nArr[1])

	rating := map[string]float32{"C":5, "GO":4.5, "Python":4}
	python, ok := rating["Python"]
	if ok {
		fmt.Println("Python:", python)
	} else {
		fmt.Println("Has no python in map.")
	}

	numbers := make(map[string]int);
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["nine"] = 9
	fmt.Println("numbers[2]:",numbers["nine"])
}
