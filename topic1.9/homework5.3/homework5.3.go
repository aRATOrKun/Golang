package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) float64 {
	b.Balance += amount
	return b.Balance
}

func (b *BankAccount) Withdraw(amount float64) float64 {
	if b.Balance-amount < 0 {
		fmt.Println("Недостачно средств")
		return b.Balance
	}

	b.Balance -= amount

	return b.Balance
}

func main() {
	bankAccount := BankAccount{"Михаил", 30000}

	bankAccount.Deposit(5000)
	fmt.Printf("Счет после пополнения: %f\n", bankAccount.Balance)

	bankAccount.Withdraw(2500)
	fmt.Printf("Счет после снятия: %f\n", bankAccount.Balance)

	bankAccount.Withdraw(50000)
	fmt.Printf("Итоговый счет: %f\n", bankAccount.Balance)
}
