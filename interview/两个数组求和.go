package interview

func add(arr1 []int, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	maxLen := max(l1, l2) + 1
	res := make([]int, maxLen)

	carry := 0
	for i := 0; i < maxLen; i++ {
		num1 := getNum(arr1, l1-1-i)
		num2 := getNum(arr2, l2-1-i)

		sum := num1 + num2 + carry

		res[maxLen-1-i] = sum % 10
		carry = sum / 10
	}

	if res[0] == 0 {
		return res[1:]
	}

	return res
}

func getNum(arr []int, index int) int {
	if index >= 0 && index < len(arr) {
		return arr[index]
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
