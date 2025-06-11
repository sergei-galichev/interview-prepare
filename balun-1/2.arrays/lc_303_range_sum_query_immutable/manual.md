[303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/)

```go
package main

type NumArray struct {
	prefixSum []int
}

/*
time: O(n), n is array length
mem: O(n) n is array length
*/

func Constructor(nums []int) NumArray {
	prefixSum := []int{0}

	for _, num := range nums {
		prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+num)
	}

	return NumArray{
		prefixSum: prefixSum,
	}
}

/*
time: O(1), substracting two number
mem: O(1) substracting two number
*/

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

```

***Оценка по времени:***

для конструктора - `O(n)`, где `n` - количество элементов слайса. для выборки -  `O(1)`, т.к. это разность двух значений

*Объяснения:* для конструктора делаем 1 проход по всем элементам, который образует `n` итераций. для выборки фактически 
выбираем 2 элемента по индексам

***Оценка по памяти:***

для конструктора - `O(n)`, где `n` - количество элементов слайса. для выборки -  `O(1)`, т.к. это разность двух значений

*Объяснения:* для конструктора делаем 1 проход по всем элементам, который образует `n` итераций. для выборки фактически
выбираем 2 элемента по индексам

**Описание решения**

Использование префиксного массива: это массив поверх основного, элементы которого это сумм предыдущих в основном.
В конструкторе создаем префиксный слайс с одним элементом "0". Идем по слайсу nums и вставляем в префиксный слайс сумму
текущего элемента и предыдущих (конкретно последнего, т.к. он является суммой предыдущих). Выборка суммы по диапазону
это разность элементов в префиксном слайсе: (right+1) - left


