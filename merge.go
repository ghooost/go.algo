package main

func Merge(arr []int) {
	if len(arr) < 2 {
		return
	}
	l := int(len(arr) / 2)
	a1 := arr[:l]
	a2 := arr[l:]
	Merge(a1)
	Merge(a2)
	buf := make([]int, len(arr))
	i1 := 0
	i2 := 0
	i3 := 0
	for {
		if i1 >= len(a1) && i2 >= len(a2) {
			break
		} else if i1 >= len(a1) {
			buf[i3] = a2[i2]
			i2++
			i3++
		} else if i2 >= len(a2) {
			buf[i3] = a1[i1]
			i1++
			i3++
		} else if a1[i1] > a2[i2] {
			buf[i3] = a2[i2]
			i2++
			i3++
		} else {
			buf[i3] = a1[i1]
			i1++
			i3++
		}
	}
	for i := 0; i < len(buf); i++ {
		arr[i] = buf[i]
	}
}
