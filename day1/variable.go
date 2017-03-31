package main

import (
	"fmt"
)

func main() {
	//定义一个变量
	//var variableName type
	//定义多个变量
	//var vname1,vname2,vname3 type
	//定义变量并初始化
	//var variableName type = value
	//同时初始化多个变量
	//var vname1,vname2,vname3 type = v1,v2,v3
	//忽略类型
	//var vname1,vname2,vname3 = v1,v2,v3
	//:=取代了var 但只能在函数内部使用，及局部变量。var用来定义全局变量
	//vname1,vname2,vname3 := v1,v2,v3

	//_(下划线)是特殊的变量名，赋予它的值会被丢弃
	//_,b := 12,13 //12被丢弃

	var1, var2, var3 := 10, 11, 12
	fmt.Printf("var1:%d,var2:%d,var3:%d", var1, var2, var3)
}