// Исправьте код так, чтобы:
// 1. Реализации интерфейса работали корректно.
// 2. Программа завершалась успешно и без паники.
// 3. Лишние вызовы методов были устранены.
// 4. Вы смогли объяснить, где были ошибки и почему они возникали.

package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) error
	Verify(amount float64) bool
}

type CreditCardProcessor struct {
	limit float64
}

func (c *CreditCardProcessor) Process(amount float64) error {
	if amount > c.limit {
		return fmt.Errorf("credit card limit exceeded")
	}

	c.limit -= amount

	fmt.Printf("Processed payment of $%.2f using CreditCard\n", amount)

	return nil
}

func (c CreditCardProcessor) Verify(amount float64) bool {
	return amount <= c.limit
}

type PayPalProcessor struct {
	balance float64
}

func (c *PayPalProcessor) Process(amount float64) error {
	if amount > c.balance {
		return fmt.Errorf("not enough balance in PayPal")
	}

	c.balance -= amount

	fmt.Printf("Processed payment of $%.2f using PayPal\n", amount)

	return nil
}

func (c *PayPalProcessor) Verify(amount float64) bool {
	return amount <= c.balance
}

func ExecutePayment(payment PaymentProcessor, amount float64) {
	if payment.Verify(amount) {
		err := payment.Process(amount)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Verification failed for amount: ", amount)
	}
}

func main() {
	creditCard := &CreditCardProcessor{limit: 100.0}
	payPal := &PayPalProcessor{balance: 200.0}

	ExecutePayment(creditCard, 50.0)
	ExecutePayment(creditCard, 50.0)
	ExecutePayment(payPal, 150.0)
	ExecutePayment(payPal, 150.0)
	ExecutePayment(payPal, 150.0)
}
