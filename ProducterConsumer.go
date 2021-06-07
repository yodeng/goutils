package main

import "fmt"

func Producer(ch1 chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch1 <- i
		fmt.Println("生产一个数据：", i)
	}
	close(ch1)
}

func Consumer(ch1 <-chan int, exch1 chan bool) {
	for v := range ch1 {
		fmt.Println("消费一个数据:", v)
	}
	exch1 <- true
}

func main() {
	ch1 := make(chan int, 10)
	exch1 := make(chan bool, 1)
	var n int = 100
	go Producer(ch1, n)
	go Consumer(ch1, exch1)
	<-exch1
	fmt.Println("运行结束")
}
