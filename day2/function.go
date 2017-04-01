package main

import (
	"fmt"
)

type testInt func(int) bool

func main() {
	//func funcName(input1 type1,input2 type2) (output1 value1,output2 value2) {
	//	//这里是处理逻辑
	//	//返回多个值
	//	return value1, value2
	a, b := 2, 3
	x, y := func1(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, x)
	fmt.Printf("%d * %d = %d\n", a, b, y)

	//变餐
	// func funcName(arg ...int)
	func2(1, 2, 3)

	// defer 当函数执行到最后 defer语句会逆序执行
	defer1()

	//函数作为值、类型
	slice := []int{1,2,3,4,5,6,7}
	odd := filter(slice, isOdd)
	even := filter(slice, isEven)
	fmt.Println("odd slice are: ", odd)
	fmt.Println("even slice are: ", even)
}

func func1(a, b int) (int, int) {
	return a + b, a * b
}
func func2(arg ...int) {
	for _, n := range arg {
		fmt.Printf("num is : %d;\t", n)
	}
	fmt.Println()
}
func defer1() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d,", i)
	}
	fmt.Println()
}

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false;
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int{
	var result []int
	for _,value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}