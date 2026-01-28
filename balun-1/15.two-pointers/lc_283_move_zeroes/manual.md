[283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)

```go
package main

func moveZeroes(nums []int) {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
```

***Оценка по времени:*** `O(n)`

*Объяснения:* 1 проход по всему слайсу

***Оценка по памяти:*** `O(1)`

*Объяснения:* по условию, нужно сместить нули in-place. 

**Описание решения**

Создаем 2 указателя slow и fast, которые равны 0. В цикле идем, пока fast меньше длины слайса с числами. Если элемент
слайса по индексу fast не равен 0, то меняем значения элементов слайса по индексам slow и fast местами. 