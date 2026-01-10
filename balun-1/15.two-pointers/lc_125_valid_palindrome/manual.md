[125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

```go
package main

func isPalindrome(s string) bool {
	first := 0
	second := len(s) - 1

	for first < second {
		if !isAlphanumeric(s[first]) {
			first++
		} else if !isAlphanumeric(s[second]) {
			second--
		} else if toLower(s[first]) == toLower(s[second]) {
			first++
			second--
		} else {
			return false
		}
	}

	return true
}

func isAlphanumeric(symbol byte) bool {
	return symbol >= 'A' && symbol <= 'Z' ||
		symbol >= 'a' && symbol <= 'z' ||
		symbol >= '0' && symbol <= '9'
}

func toLower(symbol byte) byte {
	if symbol >= 'A' && symbol <= 'Z' {
		return symbol + ('a' - 'A')
	}

	return symbol
}
```

***Оценка по времени:*** `O(n)`

*Объяснения:* проход 1 раз по всем символам в строке

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительной памяти не выделяется

**Описание решения**

создаем 2 указателя (индекса) на первый (first) и последний (second) элемент в строке. идем в цикле и проверяем,
являются ли символы под этими индексами буквенно-числовыми. затем, сравниваем в нижнем регистре равны ли символы. 