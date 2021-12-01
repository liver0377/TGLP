package reverse

const (
	len int = 10
)

func Reverse (p *[len]int) {
	for i , j := 0, len - 1; i < j; i, j = i + 1, j - 1 {
		p[i], p[j] = p[j], p[i]
	}
}

