package code

import "sync"

type BankAccount struct {
	RWMutex  sync.RWMutex
	Ballance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Ballance = account.Ballance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Ballance
	account.RWMutex.RUnlock()
	return balance
}
