package code

import (
	"fmt"
	"time"
)

func Goroutin() {
	fmt.Println("hello world")
}

func DisplayNum(number int) {
	fmt.Println("display", number)
}

func GiveMeRespone(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "bobby pratama"
}

// hanya bisa menerima data ke dalam channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "bobby pratama"
}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}
