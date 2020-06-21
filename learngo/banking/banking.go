package banking

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

func MakeAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) ChangeBalance(balance int) {
	a.balance += balance
}

func (a *Account) Withdraw(money int) error {
	if a.balance < money {
		return errors.New("you cannot withdraw ")
	}
	a.balance -= money
	return nil
}

func (a Account) CheckBalance() int {
	return a.balance
}
