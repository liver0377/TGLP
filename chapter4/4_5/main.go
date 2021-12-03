package main

import "fmt"

//return the new length
func rmAdjStr (str []string) int{
	var s = 0
	for f := 1; f < len(str); f++ {
		if str[f] != str[s] {
			s++
			str[s] = str[f]
		}
	}
	return s + 1
}

func main () {
	var str = []string{"hello", "hello", "hooooo", "hooooo", "hiiii", "good", "bye"}
	len := rmAdjStr(str)
	for i := 0; i < len; i++ {
		fmt.Println(str[i])
	}
}
