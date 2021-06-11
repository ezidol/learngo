package main

import (
	"fmt"

	"github.com/ezidol/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Peter")
	account.Deposit(100)
	fmt.Println(account.Balance())
	var err = account.Withdraw((150))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
