package main

import "fmt"

func CreateDate(ch1 chan<- int, n []int) {
	for i := 0; i < cap(n); i++ {
		ch1 <- i
	}
	close(ch1)
	fmt.Println("任务队列创建完成")
}

func Worker(ch1, ch2 chan int, exch1 chan bool) {
	for v := range ch1 {
		res := v * v // 任务执行代码块, 以求平方为例
		ch2 <- res
	}
	exch1 <- true
	fmt.Println("完成一个协程")
}

func main() {
	jobs := make([]int, 100) // 创建100个任务
	ch1 := make(chan int, 10)
	res := make(chan int, 5)
	exch := make(chan bool, 1)
	go CreateDate(ch1, jobs)

	pool := 4
	for i := 1; i <= pool; i++ { // 开启4个协程
		go Worker(ch1, res, exch)
	}

	go func() {
		for i := 1; i <= pool; i++ { // 主线程不阻塞，开启一个协程等待协程池任务完成，并关闭结果通道,
			<-exch
		}
		close(res)
	}()

	/*
	    for i := 1; i <= pool; i++ { // 若不新开启匿名函数groutine而直接用主groutine, 则后续res通道无法读, worker中res写满阻塞时，exch的写也会阻塞，此处的exch读也会阻塞，最终全部groutine, deadlock
			<-exch
		}
		close(res)
	*/

	for n := range res { // 主线程取值直到协程池全部执行完成
		fmt.Println("运行结果返回值是:", n)
	}
}
