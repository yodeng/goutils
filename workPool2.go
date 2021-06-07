package main

import "fmt"

func CreateDate(ch1 chan<- int, n []int) {
	for i := 0; i < cap(n); i++ {
		ch1 <- i
	}
	close(ch1)
	fmt.Println("任务队列创建完成")
}

func Worker(ch1, ch2 chan int) {
	for v := range ch1 {
		res := v * v // 任务执行代码块, 以求平方为例
		ch2 <- res
	}
	fmt.Println("完成一个协程")
}

func main() {
	jobs := make([]int, 100) // 创建100个任务
	ch1 := make(chan int, 10)
	res := make(chan int, 5)
	go CreateDate(ch1, jobs)

	pool := 4
	for i := 1; i <= pool; i++ { // 开启4个协程
		go Worker(ch1, res)
	}

	for i := 0; i < cap(jobs); i++ { // 主线程不阻塞, 等待结果输出
		n := <-res
		fmt.Println("运行结果返回值是:", n)
	}
	fmt.Println("运行结束")
}
