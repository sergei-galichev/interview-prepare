[977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/)

```go
package main

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1

	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		lhs := abs(nums[left])  // math.Abs
		rhs := abs(nums[right]) // math.Abs

		if lhs >= rhs {
			result[i] = lhs * lhs
			left++
		} else {
			result[i] = rhs * rhs
			right--
		}
	}

	return result
}

func abs(number int) int {
	if number < 0 {
		return -number
	} else {
		return number
	}
}
```

***Оценка по времени:*** `O(n)`

*Объяснения:* проходим по массиву 1 раз

***Оценка по памяти:*** `O(1)`

*Объяснения:* если не считать результирующий слайс

**Описание решения**

Создаем 2 индекса на начало и конец слайса. В цикле идем с конца слайса. Берем абсолютные значения элементов и 
сравниваем. Если левый больше либо равен правому, то левый возводим в квадрат и добавляем в результирующий слайс. Левый
индекс увеличиваем на 1. Иначе добавляем в слайс квадрат правого и уменьшаем на 1 правый индекс