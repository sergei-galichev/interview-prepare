[344. Reverse String](https://leetcode.com/problems/reverse-string/)

```go
package main

func reverseString(s []byte) {
	first := 0
	second := len(s) - 1

	for first < second {
		s[first], s[second] = s[second], s[first]

		first++
		second--
	}
}
```

***Оценка по времени:*** `O(n)`

*Объяснения:* 1 проход не более, чем по всем элементам

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительной памяти не выделяется

**Описание решения**

создаем 2 указателя (индекса) на первый (first) и последний (second) элемент в слайсе. идем в цикле и меняем значения 
под индексами местами. сдвигаем first вперед, second - назад. цикл завершится, если first >= second