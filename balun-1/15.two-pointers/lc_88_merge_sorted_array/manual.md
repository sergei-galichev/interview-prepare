[88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)

```go
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	first := m - 1
	second := n - 1

	for idx := m + n - 1; idx >= 0; idx-- {
		if second < 0 {
			break
		}

		if first >= 0 && nums1[first] >= nums2[second] {
			nums1[idx] = nums1[first]

			first--
		} else {
			nums1[idx] = nums2[second]

			second--
		}
	}
}
```

***Оценка по времени:*** `O(m + n)`

*Объяснения:* проход по элементам обоих слайсов

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительная память не используется

**Описание решения**

создаем 2 указателя (индекса) на последние элементы слайсов: `first := m - 1` и `second := n - 1`. в цикле начиная с
конца 1-го слайса `idx := m + n - 1`. проверяем, если индекс second < 0, то выход, т.к. уже смержили все элементы. сравниваем, если индекс 
first валидный и элемент в 1-м слайсе >= элементу во 2-м слайсе, то вставляем его по индексу idx в 1-й слайс. иначе,
вставляем элемент 2-го слайса.