package account

import "sync"

// Define the Account type here.
type Account struct {
	actived bool
	balance int64
	sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		actived: true,
		balance: amount,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	if !a.actived {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.actived {
		return 0, false
	}

	if a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.actived {
		return 0, false
	}

	balance := a.balance
	a.actived = false
	a.balance = 0

	return balance, true
}
