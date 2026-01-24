[349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)

```go
package main

import (
	"slices"
)

func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	idx1 := 0
	idx2 := 0

	var result []int

	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			result = append(result, nums1[idx1])

			for idx1 < len(nums1) && nums1[idx1] == result[len(result)-1] {
				idx1++
			}

			for idx2 < len(nums2) && nums2[idx2] == result[len(result)-1] {
				idx2++
			}
		}
	}

	return result
}
```

***Оценка по времени:*** `O(n x log(n) + m x log(m))`

*Объяснения:* если используется быстрая сортировка, то сложность логарифмическая

***Оценка по памяти:*** `O(1)`

*Объяснения:* если не считать результирующий слайс

**Описание решения**

Сортируем входные слайсы. Устанавливаем 2 индекса в начальные положения каждого слайса. Далее, пока оба индекса не за
пределами длин слайсов, идем в цикле. Сравниваем, если элемент 1-го слайса меньше элемента 2-го, то двигаем индекс 1-го.
Если наоборот, то инкрементируем индекс 2-го. Иначе добавляем в результирующий слайс, т.к. они равны. Также смещаем 
индексы, если следующий элемент равен предыдущему и не вышли за пределы слайса 