package code

import (
	"fmt"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bobby",
		Balance: 100000,
	}

	user2 := UserBalance{
		Name:    "Pratama",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 15000)

	time.Sleep(5 * time.Second)

	fmt.Println("user :", user1.Name, "ballance :", user1.Balance)
	fmt.Println("user :", user2.Name, "ballance :", user2.Balance)
}
