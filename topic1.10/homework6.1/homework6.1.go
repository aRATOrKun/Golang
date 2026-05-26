package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) string
}

func (cc CreditCard) Process(amount float64) string {
	if cc.Balance-amount >= 0 {
		return "Успешный платеж рублями"
	} else {
		return "Недостаточно средств"
	}

}

func (cw CryptoWallet) Process(amount float64) string {
	btcToRub := cw.BalanceBTC * 5507354
	ethToRub := cw.BalanceETH * 151223

	if btcToRub-amount >= 0 || ethToRub-amount >= 0 {
		return "Успешный платеж криптовалютой"
	} else {
		return "Недостаточно средств"
	}
}

type CreditCard struct {
	Owner   string
	Balance float64
}

type CryptoWallet struct {
	OwnerID    int
	BalanceBTC float64
	BalanceETH float64
}

func main() {
	p1 := CreditCard{"Михаил", 100000}
	p2 := CreditCard{"Алексей", 20000}
	p3 := CryptoWallet{6735912, 0, 0.1}
	p4 := CryptoWallet{1235342, 0.25, 109}

	users := []PaymentProcessor{p1, p2, p3, p4}

	for _, user := range users {
		fmt.Println(user.Process(100000))
	}
}
