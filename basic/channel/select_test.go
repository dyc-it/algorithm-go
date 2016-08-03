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
		time.Sleep(time.Duration(n+1) * time.Second)
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
















