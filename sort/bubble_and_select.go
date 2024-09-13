package sort

// 冒泡和选择本质上没有多大区别，都是选出最小数，然后放置于合适的位置

// 冒泡排序
// 每一轮循环，都将选出最小的数，将最小的数置于数组最前方或者最后方，类似于泡泡，一个一个往上
// 关键在于，每次循环都会选出最小的数，即第二层的循环次数会一次次的变少
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 选择排序
// 每一轮循环，选出一个最小的数，将其置于合适的的位置
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		tmp := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[tmp] {
				tmp = j
			}
		}
		arr[i], arr[tmp] = arr[tmp], arr[i]
	}
}
