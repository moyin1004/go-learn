package design

import (
	"fmt"
	"sync"
)

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func isNum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 1 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for idx, c := range cs {
		go func(c <-chan int, idx int) {
			for n := range c {
				fmt.Println(idx, n)
				out <- n
			}
			wg.Done()
		}(c, idx)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func TestPipeline() {
	nums := makeRange(1, 100)
	in := echo(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	// 数组分块计算
	for i := range chans {
		chans[i] = sum(isNum(in))
	}

	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}
