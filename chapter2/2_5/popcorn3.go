package popcount


func PopCount3(x uint64) int {
	a := 0
	for i := 0; i < 64; i++ {
		if x != 0 {
			x = x & (x - 1)
			a++
		} else {
			break;
		}
	}
	return a
}
