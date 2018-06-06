package account

import (
	"sync"
)

type empty interface{}

// Account doc here
type Account struct {
	balance int64
	closed  bool
	mutex   sync.Mutex
}

// Open doc here
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close doc here
func (account *Account) Close() (payout int64, ok bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()

	if account.closed {
		return 0, false
	}
	account.closed = true
	return account.balance, true
}

// Balance doc here
func (account *Account) Balance() (balance int64, ok bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()

	if account.closed {
		return account.balance, false
	}
	return account.balance, true
}

// Deposit doc here
func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()

	if account.closed {
		return account.balance, false
	}
	if account.balance+amount < 0 {
		return account.balance, false
	}
	account.balance += amount
	return account.balance, true
}
