package sort

func quickSort(arr []int, low, high int) {
	if low < high {
		mid := partition(arr, low, high)
		quickSort(arr, low, mid-1)
		quickSort(arr, mid+1, high)
	}

}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	// 将小于基准值的数，都放到一边，大于的数放到另一边
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
