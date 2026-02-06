package main

func main() {

}

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
