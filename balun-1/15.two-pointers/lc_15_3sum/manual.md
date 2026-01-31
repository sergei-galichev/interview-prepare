[15. 3Sum](https://leetcode.com/problems/3sum/)

```go
package main

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var result [][]int

	for idx := 0; idx < len(nums)-2; idx++ {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		left := idx + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[idx] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[idx], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}
```

***Оценка по времени:*** `O(n x log(n) + n^2)` или `O(n^2)` 

*Объяснения:* быстрая сортировка и двойной проход по всем элементам слайса

***Оценка по памяти:*** `O(1)`

*Объяснения:* если не считать результирующий слайс, то дополнительной памяти не выделяется

**Описание решения**

Сначала сортируем слайс. Далее, создаем индекс idx и в цикле идем по слайсу с  до 2-го элемента с конца. Проверка: не
равно ли текущее значение предыдущему. Создаем индексы: left = idx + 1, right = длина слайса - 1. Во 2-м цикле идем пока
left < right. Находим сумму трех элементов по индексам idx, left, right и сравниваем с 0. Если сумма меньше 0, то 
двигаем left на 1 вперед. Если же сумма больше 0, то двигаем right на 1 назад. В ином случае, добавляем в результирующий
2-мерный слайс новый слайс с элементами по индексам idx, left, right. Двигаем индексы left вперед и right назад.
Проверяем, что left < right и текущий элемент с индексом left не равен предыдущему. Иначе двигаем left вперед.
Проверяем, что left < right и текущий элемент с индексом right не равен следующему. Иначе двигаем right назад.