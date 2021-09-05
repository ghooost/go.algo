package quick

func Quick(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[0]
	iL := 1
	iR := len(arr) - 1
	for {
		if iL >= iR {
			if pivot > arr[iR] {
				arr[0], arr[iR] = arr[iR], arr[0]
			}
			break
		}
		for arr[iL] < pivot {
			iL++
		}
		for arr[iR] > pivot {
			iR--
		}
		if iL < iR {
			arr[iL], arr[iR] = arr[iR], arr[iL]
		}
	}
	if iR > 0 {
		Quick(arr[:iR])
	}
	if iR < len(arr)-1 {
		Quick(arr[iR+1:])
	}
}
