package code

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestGoroutin(t *testing.T) {
	go Goroutin()
	fmt.Println("pp")

	time.Sleep(1 * time.Second)
}

func TestDisplayNum(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNum(i)
	}
	time.Sleep(10 * time.Second)
}

func TestChannel(t *testing.T) {
	channel := make(chan string)

	go GiveMeRespone(channel)

	data := <-channel
	fmt.Println(data)
	defer close(channel)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	defer close(channel)
}

// mentesting lock go routine menggugnakan mutex
func TestMutex(t *testing.T) {
	var x = 0

	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter :", x)
}

//chanel untuk melihat range chanel

func TestRangeChanel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "data perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done")
}
