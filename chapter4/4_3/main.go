package main

import (
	"awesomeProject/reverse"
	"fmt"
)
func main () {
	var a = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse.Reverse(&a)
	fmt.Println(a)
}
