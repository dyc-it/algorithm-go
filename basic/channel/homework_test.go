package channel

import (
	"fmt"
	"testing"
	"time"
)

const STUDENT_NUM = 10

type Homework struct {
	num int
}

func student(hwChan chan Homework, num int) {
	hwChan <- Homework{num:num}
}

func teacher(hwChan chan Homework) {
	for i := 0; i < STUDENT_NUM; i++ {
		hw := <-hwChan
		fmt.Println(hw)
	}
}

func TestHomework1(t *testing.T) {
	hwChan := make(chan Homework, STUDENT_NUM)

	for i := 0; i < STUDENT_NUM; i++ {
		go student(hwChan, i)
	}

	go teacher(hwChan)

	time.Sleep(3 * time.Second)
}







