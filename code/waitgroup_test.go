package code

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSyncWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunSync(group)
	}
	group.Wait()
	fmt.Println("done")
}

// testing once waitgroup func akan di eksekusi hanya 1 kali

func TestOnce(t *testing.T) {
	var once sync.Once
	var grouping sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			grouping.Add(1)
			once.Do(OnlyOnce)
			grouping.Done()
		}()
	}

	grouping.Wait()
	fmt.Println(counter)
}

// sync.Pool sync untuk mengambil dan mengembalikan data yang sudah ada di go routin
func TestSyncPool(t *testing.T) {
	var pool = sync.Pool{

		New: func() interface{} {
			return "new"
		},
	}

	pool.Put("Bobby")
	pool.Put("Pratama")
	pool.Put("Maulana")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	time.Sleep(3 * time.Second)
}

// sync.Map untuk membuat map di goroutin agar tidak terjadi deadlock atau race condition
func TestSyncMap(t *testing.T) {
	var data sync.Map

	var addToMap = func(value int) {
		data.Store(value, value)
	}

	for i := 0; i < 10; i++ {
		go addToMap(i)
	}

	time.Sleep(3 * time.Second)
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", key)
		return true
	})
}

// sync.Cond Cond adalah adalah implementasi locking berbasis kondisi
func TestSynCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait()
}

//atomic sync adalah Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent

func TestAtomic(t *testing.T) {
	var group sync.WaitGroup

	var counter int64 = 0

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter", counter)
}
