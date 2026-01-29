[167. Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

```go
package main

func twoSum(numbers []int, target int) []int {
	var result []int

	left := 0
	right := len(numbers) - 1

	for left != right {
		if numbers[left]+numbers[right] == target {
			result = append(result, left+1, right+1)
			break
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}
	}

	return result
}
```

***Оценка по времени:*** `O(n)`

*Объяснения:* проход 1 раз по слайсу

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительная память не выделяется

**Описание решения**

Создаем 2 указателя на начало (left) и конец (right) слайса. В цикле идем по слайсу пока индексы left != right. Если
сумма элементов по индексам left и right равна заданной, то добавляем их в результирующий слайс, добавив 1 к каждому.
Если сумма элементов больше заданной, то уменьшаем right на 1, если меньше, уменьшаем left на 1. 