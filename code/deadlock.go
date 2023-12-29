package code

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (User *UserBalance) Lock() {
	User.Mutex.Lock()
}

func (User *UserBalance) Unlock() {
	User.Mutex.Unlock()
}

func (User *UserBalance) Change(amount int) {
	User.Balance = User.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("lock", user1.Name)
	user1.Change(amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
