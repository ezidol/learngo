package main

import (
	"fmt"

	"github.com/ezidol/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Peter")
	account.Deposit(100)
	fmt.Println(account.Balance())
}
