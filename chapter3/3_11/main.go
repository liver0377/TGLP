package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	i := strings.LastIndex(s, ".")
	var  suffix string
	var prefix string
	if i != -1 {
		suffix = s[i:]
		prefix = s[:i]
	}
	var n = len(prefix)
	var sign byte
	if n > 0 && (s[0] == '+' || s[0] == '-') {
		sign = s[0]
		n--
		prefix = prefix[1:]
	}
	if n <= 3 {
		return s
	}
	if sign == 0 {
		return commaInt(prefix) + suffix
	} else {
		return string(sign) + commaInt(prefix) + suffix
	}

}

func commaInt(s string) string {
	var n int = len(s)

	if n <= 3 {
		return  s
	}

	return commaInt(s[: n - 3]) + "," + s[n - 3:]
}
func main() {
	var s = comma("+12346.1234")
	fmt.Println(s)
}
