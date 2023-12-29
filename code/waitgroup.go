package code

import (
	"fmt"
	"sync"
	"time"
)

// harus menggunakan done untuk menghentikan func go routin bahwa func selesai di jalankan
func RunSync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}

var counter = 0

func OnlyOnce() {
	counter++
}

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println("done")
	cond.L.Unlock()
	group.Done()
}
