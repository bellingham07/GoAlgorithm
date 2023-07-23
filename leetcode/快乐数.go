package leetcode

func isHappy(n int) bool {
	getNum := func(num int) int {
		sum := 0
		for num > 0 {
			sum += (num % 10) * (num % 10)
			num /= 10
		}
		return sum
	}

	m1 := map[int]bool{}
	for n != 1 && !m1[n] {
		n, m1[n] = getNum(n), true
	}
	return n == 1
}
