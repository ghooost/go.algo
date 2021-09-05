package bubble

func Bubble(arr []int) int {
	nIter := 0
	for {
		nIter++
		changes := 0
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				changes++
			}
		}
		if changes == 0 {
			return nIter
		}
	}
}
