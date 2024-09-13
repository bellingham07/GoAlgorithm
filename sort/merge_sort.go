package sort

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	res := merge(left, right)
	return res
}

func merge(left, right []int) []int {
	i, j := 0, 0
	l, r := len(left), len(right)

	var res []int

	for {
		if i >= l || j >= r {
			break
		}

		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	// 特别注意这里，不是数组的长度，而是指针的值
	if i != l {
		res = append(res, left[i:]...)
	} else {
		res = append(res, right[j:]...)
	}

	return res
}
