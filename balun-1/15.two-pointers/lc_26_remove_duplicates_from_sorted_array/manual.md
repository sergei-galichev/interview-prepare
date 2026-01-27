[26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/

```go
package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 1

	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
	}

	return slow + 1
}
```

***Оценка по времени:*** `O(n)`

*Объяснения:* 1 проход по всем элементам слайса

***Оценка по памяти:*** `O(1)`

*Объяснения:* по условию нужно in-place решение. дополнительной памяти не выделяется

**Описание решения**

Corner case: если длина слайса 0, то вернем 0. Создаем два указателя на 0-й (slow) и 1-й (fast) элементы. В цикле идем
по слайсу, пока fast меньше длины слайса. Если значения элементов по slow и fast индексам равны, то увеличиваем fast
на 1. Если не равны, то увеличиваем slow на 1, элементу слайса по индексу slow присваиваем значение по индексу fast и
увеличиваем fast на 1. Функция возвращает slow + 1, т.к. индексация в слайсе начинается с 0 