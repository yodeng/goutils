package main

import "fmt"
import "sync"

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func Producer(ch1 chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch1 <- i
		fmt.Println("生产一个数据：", i)
	}
	close(ch1)
}

func Consumer(ch1 <-chan int, wg *sync.WaitGroup) {
	for v := range ch1 {
		fmt.Println("消费一个数据:", v)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	w := worker{
		in: make(chan int, 10),
		wg: &wg,
	}
	var n int = 100
	go Producer(w.in, n)
	go Consumer(w.in, &wg)
	wg.Wait()
	fmt.Println("运行结束")
}
