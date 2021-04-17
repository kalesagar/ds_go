package programs

//LargeNumber ...
func LargeNumber(arr []int) int {
	arrLength := len(arr)
	pivot := arrLength / 2
	for (pivot+1 < arrLength) && (pivot > 0) {
		if arr[pivot] > arr[pivot+1] && arr[pivot] > arr[pivot-1] {
			break
		}
		if arr[pivot] < arr[pivot+1] {
			pivot = pivot + 1
		} else {
			pivot = pivot - 1
		}
	}
	return arr[pivot]
}
