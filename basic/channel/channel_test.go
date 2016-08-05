package channel

import (
	"testing"
	"fmt"
	"time"
	"runtime"
)

func TestChannel01(t *testing.T) {
	go func() {
		fmt.Println("Message From goroutine")
	}()
}

func TestChannel02(t *testing.T) {
	go func() {
		fmt.Println("Message From goroutine")
	}()
	time.Sleep(time.Second)
}

func TestChannel03(t *testing.T) {
	var ch chan int
	ch <- 1
	fmt.Println("goroutine channel test")
}

func TestChannel04(t *testing.T) {
	var ch chan int
	<-ch
	fmt.Println("goroutine channel test")
}

func TestChannel05(t *testing.T) {
	var ch chan int
	close(ch)
	fmt.Println("goroutine channel test")
}

func TestChannel06(t *testing.T) {
	ch := make(chan int)
	close(ch)
	ch <- 1
	fmt.Println(<-ch)
}

func TestChannel07(t *testing.T) {
	ch := make(chan int)
	close(ch)
	<-ch
	fmt.Println(<-ch)
}

/*
读无缓冲的channel会一直阻塞线程,直到有线程写入了channel,所以下面的代码会造成死锁
 */
func TestChannel08(t *testing.T) {
	ch := make(chan int)
	<-ch
}

/*
写无缓冲的channel会一直阻塞线程,直到有线程读了channel,所以下面的代码会造成死锁
 */
func TestChannel09(t *testing.T) {
	ch := make(chan int)
	ch <- 1
}

func TestChannel10(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int)
	go func() {
		fmt.Println("Before write to channel")
		ch <- 1
		fmt.Println("After write to channel")
	}()
	<-ch
}

func TestChannel11(t *testing.T) {
	ch := make(chan int)
	go func() {
		fmt.Println("Before write to channel")
		// 写入后主线程就不会阻塞了
		ch <- 1
		time.Sleep(time.Millisecond)
		fmt.Println("After write to channel")
	}()
	<-ch
}

func TestChannel12(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int)
	go func() {
		fmt.Println("goroutine start")
		ch <- 1
		fmt.Println("goroutine end")
	}()
	fmt.Println("main thread start")
	<-ch
	fmt.Println("main thread end")
}

/*
Q:
1.在Go语言1.5之前的版本中,执行如下的代码,两个loop函数的输出会有先后顺序吗?
1.在Go语言1.5及其之后的版本中,执行如下的代码,两个loop函数的输出会有先后顺序吗?

A:
1.会先执行完一个loop,再执行另一个
2.看CPU的核数,如果CPU是多核的,两个loop函数是同时输出的
解析:在Go1.5之前的版本中,默认使用一个核;在1.5及之后的版本中,默认使用所有可用核。
 */
func TestChannel13(t *testing.T) {
	var quit chan int = make(chan int)

	loop := func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
		}
		quit <- 0
	}

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
	fmt.Println()
}


/*
Q:下面这段代码的输出是?
A:没有输出,因为只使用了一个核,for死循环占据了单核的所有资源,导致goroutine没有机会执行
 */
func TestChannel14(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("Hello")
	}()
	for ; ; {

	}
}

/*
Q:下面这段代码的输出是?
A:Hello
 */
func TestChannel15(t *testing.T) {
	runtime.GOMAXPROCS(2)
	go func() {
		fmt.Println("Hello")
	}()
	for ; ; {

	}
}


/*
Q:ret1-4和ok1-4的值分别是?
A:
ret1=1, ok1=true;因为channel使用FIFO的原则,未关闭时,ok返回true
ret2=2, ok2=true;channel虽然关闭了,但是缓存中仍然有数据
ret3=3, ok3=true;channel虽然关闭了,但是缓存中仍然有数据
ret4=0, ok4=false;ok为false表示channel被关闭并且缓冲区中没有数据了
 */

func TestChannel16(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	ret1, ok1 := <-ch
	fmt.Println(ret1, ok1)

	close(ch)

	ret2, ok2 := <-ch
	fmt.Println(ret2, ok2)

	ret3, ok3 := <-ch
	fmt.Println(ret3, ok3)

	ret4, ok4 := <-ch
	fmt.Println(ret4, ok4)
}


/*
Q:下面程序的输出是?
A:死锁
因为channel的缓存大小是3,当缓冲区写满后,如果再往缓冲区写数据,会一直阻塞
 */
func TestChannel17(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
}


/*
Q:下面程序的输出是?
A:死锁
因为channel的缓存大小是3,没有任何写入;从一个没有数据的channel中读数据,会一直阻塞
 */
func TestChannel18(t *testing.T) {
	ch := make(chan int, 3)
	<-ch
}

/*
Q:下面的程序会阻塞吗?
A:不会,关闭channel,对无缓冲的channel,关闭时,也不会阻塞
 */
func TestChannel19(t *testing.T) {
	c := make(chan bool)
	go func(){
		close(c)
	}()
	<-c
}




