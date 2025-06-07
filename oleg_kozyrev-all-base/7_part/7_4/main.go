/*
Вы разрабатываете систему для обработки финансовых транзакций.
Каждая транзакция проходит несколько этапов обработки:

1. Чтение транзакций из исходных данных.
2. Фильтрация транзакций: убираем транзакции с отрицательными суммами.
3. Конвертация валюты: преобразуем сумму транзакции в доллары.
4. Сохранение результатов: записываем обработанные транзакции в итоговый список.
*/

package main

import (
	"fmt"
	"math/rand/v2"
)

type Transaction struct {
	ID     int64
	Amount float64
}

func generateTransaction(count int) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for i := 0; i < count; i++ {
			out <- Transaction{
				ID:     int64(i),
				Amount: rand.Float64()*200 - 100,
			}
		}

		close(out)
	}()

	return out
}

func filterTransaction(in <-chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for transaction := range in {
			fmt.Printf("---Transaction ID: %d, Amount: %.2f\n", transaction.ID, transaction.Amount)

			if transaction.Amount >= 0 {
				out <- transaction
			}
		}

		close(out)
	}()

	return out
}

func convertTransaction(in <-chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for transaction := range in {
			if transaction.Amount >= 0 {
				transaction.Amount *= 0.013
				out <- transaction
			}
		}

		close(out)
	}()

	return out
}

func saveTransaction(in <-chan Transaction) {
	for transaction := range in {
		fmt.Printf("Transaction ID: %d, Amount: %.2f USD\n", transaction.ID, transaction.Amount)
	}

}

func main() {
	transactions := generateTransaction(10)
	filteredTransactions := filterTransaction(transactions)
	convertedTransactions := convertTransaction(filteredTransactions)
	saveTransaction(convertedTransactions)
}
