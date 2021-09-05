package merge

func Merge(arr []int) {
	if len(arr) < 2 {
		return
	}
	l := int(len(arr) / 2)
	a1 := arr[:l]
	a2 := arr[l:]
	Merge(a1)
	Merge(a2)
	buf := make([]int, 0)
	i1 := 0
	i2 := 0
	for {
		if i1 >= len(a1) && i2 >= len(a2) {
			break
		} else if i1 >= len(a1) {
			buf = append(buf, a2[i2])
			i2++
		} else if i2 >= len(a2) {
			buf = append(buf, a1[i1])
			i1++
		} else if a1[i1] > a2[i2] {
			buf = append(buf, a2[i2])
			i2++
		} else {
			buf = append(buf, a1[i1])
			i1++
		}
	}
	for i := 0; i < len(buf); i++ {
		arr[i] = buf[i]
	}
}
