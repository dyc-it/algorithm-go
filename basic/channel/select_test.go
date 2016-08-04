package channel

import (
	"testing"
	"fmt"
	"time"
)


/*
Q:程序的输出是?
A:随机的0和1
 */
func TestSelect01(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println(0)
			case <-ch:
				fmt.Println(1)
			}
		}
	}()

	go func() {
		for {
			ch <- 0
		}
	}()

	time.Sleep(1 * time.Millisecond)
}



/*
以下代码实现了对ch写入超时的功能。如果在n秒内没有写入,就不会一直阻塞在ch的读操作上了。
 */
func TestTimeout(t *testing.T) {
	n := 1
	timeout := make(chan bool, 1)

	go func() {
		// 不能直接使用   n * time.Second,会包类型不匹配的编译错误。
		time.Sleep(time.Duration(n) * time.Second)
		timeout <- true
	}()

	ch := make(chan int)

	go func() {
		// 如果休眠时间小于n秒,会写入ch
		//time.Sleep(time.Duration(n-1) * time.Second)

		// 如果休眠时间大于n秒,来不及写入ch,主线程就结束了
		time.Sleep(time.Duration(n + 1) * time.Second)
		fmt.Println("before writing ch")
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("read ch occurs")
	case <-timeout:
		fmt.Println("timeout")
	}

}

/*
Q:下面代码的输出是?
A:deadlock
因为对于没有缓冲区的channel,只有读写操作都发生,才不会阻塞。
下面的代码中,程序会阻塞在ch的写操作上(ch <- 1)导致死锁
 */
func TestIsChannelFull1(t *testing.T) {
	ch := make(chan int)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("Channel if full")
	}
}

/*
Q:下面代码的输出是?
A:"Channel if full"
对于带有缓冲区的channel,写满之后,再写入,会阻塞。
所以在select语句中, ch <- 2 会阻塞,则执行default语句。
 */
func TestIsChannelFull2(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("Channel if full")
	}
}














