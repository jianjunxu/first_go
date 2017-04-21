package main

import (
	"fmt"
)

var complete chan int = make(chan int)
func main(){
	go loop()
	//loop()
	//time.Sleep(time.Second) // 停顿一秒

	//创建信道
	//var channel chan int = make(chan int)
	// 或
	//channel := make(chan int)

	//var message chan string = make(chan string)
	//go func(msg string) {
	//	message <- msg //存消息
	//}("123")
	//
	//fmt.Println(<-message) //取消息

	<- complete // 直到线程跑完, 取到消息. main在此阻塞住
}

func loop() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", i)
	}
	complete <- 0 // 执行完毕了，发个消息
}