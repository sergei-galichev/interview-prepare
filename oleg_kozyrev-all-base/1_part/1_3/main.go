package main

import "fmt"

type Account struct {
	Balance int
}

func main() {
	initialBalance := 1000
	account := &Account{
		Balance: initialBalance,
	}

	defer printBalance("Initial balance", account.Balance)
	defer printBalance("Current balance", account.Balance)
	defer printAccountBalance("Pointer to balance", account)

	account.Balance += 500
	updateBalance(account, 200)
	account = &Account{
		Balance: 300,
	}
}

func updateBalance(acc *Account, amount int) {
	acc.Balance -= amount
}

func printAccountBalance(label string, acc *Account) {
	fmt.Printf("%s: %d\n", label, acc.Balance)
}

func printBalance(label string, balance int) {
	fmt.Printf("%s: %d\n", label, balance)
}
