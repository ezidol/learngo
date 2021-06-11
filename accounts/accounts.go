package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoWithdraw = errors.New("You can't withdraw. You don't have enough balance")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: "Peter", balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoWithdraw
	}
	a.balance -= amount
	return nil
}

// ChangeOwner on your account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner on your account
func (a Account) Owner() string {
	return a.owner
}

// String on your account
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
