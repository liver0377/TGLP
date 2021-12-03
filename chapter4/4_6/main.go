package main

import (
	"fmt"
	"unicode"
)

//return the new len
func rmAdjSpace (str []byte) int{
	var s = 0
	for f := 1; f < len(str); f++ {
		if unicode.IsSpace(rune(str[s])) && unicode.IsSpace(rune(str[f])) {
			continue
		}
		s++
		str[s] = str[f]
	}
	return s + 1
}
func main ()  {
	var b []byte = []byte("xfeh            fheih")
	len := rmAdjSpace(b)
	for i := 0; i < len; i++ {
		fmt.Printf("%c", b[i])
	}
}
