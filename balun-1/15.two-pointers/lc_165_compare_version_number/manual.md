[165. Compare Version Numbers](https://leetcode.com/problems/compare-version-numbers/)

```go
package main

func compareVersion(version1 string, version2 string) int {
	left := 0
	right := 0

	for left < len(version1) || right < len(version2) {
		leftNumber := 0

		for left < len(version1) && version1[left] != '.' {
			digit := version1[left] - byte('0')
			leftNumber = leftNumber*10 + int(digit)
			left++
		}

		rightNumber := 0

		for right < len(version2) && version2[right] != '.' {
			digit := version2[right] - byte('0')
			rightNumber = rightNumber*10 + int(digit)
			right++
		}

		if leftNumber < rightNumber {
			return -1
		} else if leftNumber > rightNumber {
			return 1
		}

		left++
		right++
	}

	return 0
}
```

***Оценка по времени:*** `O(max(m,n))`

*Объяснения:* максимальная длина из двух строк, т.к. проходим по всем символам 1 раз

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительной памяти не выделяется

**Описание решения**

Создаем 2 указателя для каждой строки. В цикле делаем проход по символам пока указатели не выходят за пределы длин строк.
Идем до символа `.` либо конца каждой строки. Символы цифр парсим в число и затем их сравниваем. Если 1-е число больше
2-го, то возвращаем `1`. Если наоборот, то `-1`. Если внешний цикл завершается, то возвращаем `0`.