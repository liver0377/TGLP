package rotate


//left rotate
func Rotate (s []int, n int) {
	if n > len(s) {
		return
	}
	var temp int
	for i := 0; i < n; i++ {
		temp = s[0]
		copy(s, s[1:])
		s[len(s) - 1] = temp
	}
}

