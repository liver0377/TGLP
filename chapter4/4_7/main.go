package main

import "fmt"

//return the new len
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main ()  {
	var b []byte = []byte("xfeh            fheih")
	reverse(b)
	fmt.Printf("%s", b)
}
