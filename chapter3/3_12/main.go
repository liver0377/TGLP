package main

import "fmt"

func  isAnagma (s1, s2 string) bool{
	var n1 int = len(s1)
	var n2 int = len(s2)
	if n1 != n2 {
		return false
	}

	var num map[byte]int
	num = make(map[byte]int)

	for i := 0; i < n1; i++ {
		num[s1[i]]++
	}
	for i := 0; i < n2; i++ {
		if num[s2[i]] == 0 {
			return false
		}
	}
	return true
}
func main() {
	var s1, s2 string
	s1 = "hello world"
	s2 = "world hello"
	fmt.Println(isAnagma(s1, s2))
}
