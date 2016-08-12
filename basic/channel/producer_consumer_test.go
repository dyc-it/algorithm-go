package channel

import (
	"fmt"
	"testing"
	"time"
)

const TOTAL int = 10

func Producer(queue chan<- int) {
	for i := 0; i < TOTAL; i++ {
		queue <- i
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < TOTAL; i++ {
		v := <-queue
		fmt.Println("receive:", v)
	}
}

/*
Q:函数的输出是?
A:没有输出,因为主线程没有阻塞,直接结束
*/
func TestProducerConsumer1(t *testing.T) {
	queue := make(chan int)
	go Producer(queue)
	go Consumer(queue)
}

func TestProducerConsumer2(t *testing.T) {
	queue := make(chan int)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(2 * time.Second) //让Producer与Consumer完成
}
