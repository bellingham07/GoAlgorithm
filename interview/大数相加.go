package interview

import (
	"strconv"
)

//num1 := "123456789012345678901234567890"
//num2 := "987654321098765432109876543210"
//sum := addBigNumbers(num1, num2)
//fmt.Println("Sums:", sum)
//fmt.Println("Suma:", bigNumSum(num1, num2))

func bigNumSum(num1, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	for i := len(num2); i < len(num1); i++ {
		num2 = "0" + num2
	}

	carry := 0
	sum := ""

	//fmt.Println("num1", num1)
	//fmt.Println("num2", num2)
	for i := len(num1) - 1; i >= 0; i-- {
		digtal1, _ := strconv.Atoi(string(num1[i]))
		digtal2, _ := strconv.Atoi(string(num2[i]))

		curSum := (digtal1+digtal2)%10 + carry
		carry = (digtal1 + digtal2) / 10

		sum = strconv.Itoa(curSum) + sum
	}

	if carry > 0 {
		sum = strconv.Itoa(carry) + sum
	}

	return sum
}

func addBigNumbers(num1, num2 string) string {
	// 确保num1是较长的字符串
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	// 将两个数字字符串补齐到相同长度
	diff := len(num1) - len(num2)
	for i := 0; i < diff; i++ {
		num2 = "0" + num2
	}
	//
	//fmt.Println("num1", num1)
	//fmt.Println("num2", num2)
	// 进位标志
	carry := 0
	sum := ""

	// 从字符串的末尾逐位相加
	for i := len(num1) - 1; i >= 0; i-- {
		digit1, _ := strconv.Atoi(string(num1[i]))
		digit2, _ := strconv.Atoi(string(num2[i]))

		currentSum := digit1 + digit2 + carry

		// 处理进位
		carry = currentSum / 10
		currentSum %= 10

		// 将当前位的结果加入结果字符串
		sum = strconv.Itoa(currentSum) + sum
	}

	// 处理最后的进位
	if carry > 0 {
		sum = strconv.Itoa(carry) + sum
	}

	return sum
}
