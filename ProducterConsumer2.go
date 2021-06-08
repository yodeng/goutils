package main

import "fmt"

type worker struct {
	in   chan int
	done chan bool
}

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
	w := worker{
		in:   make(chan int, 10),
		done: make(chan bool, 1),
	}
	var n int = 100
	go Producer(w.in, n)
	go Consumer(w.in, w.done)
	<-w.done
	fmt.Println("运行结束")
}
