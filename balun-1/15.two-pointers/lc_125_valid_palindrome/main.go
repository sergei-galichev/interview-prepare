package main

func main() {

}

func isPalindrome(s string) bool {
	first := 0
	second := len(s) - 1

	for first < second {
		// unicode.IsNumber()
		// unicode.IsLetter()
		// unicode.ToLower()

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
