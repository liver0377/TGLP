package main

import (
	"awesomeProject/rotate"
	"fmt"
)
func main () {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate.Rotate(s, 3)
	fmt.Println(s)
}
