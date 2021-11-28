package popcount

func PopCount2(x uint64) int {
	a := 0
	for i:= 0; i < 64; i++ {
		if x & 1 == 1 {
			a++
		}
		x = x >> 1
	}
	return a
}
