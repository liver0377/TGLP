package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	var firstIndex int = len(s) % 3
	var buf bytes.Buffer
	buf.WriteString(s[:firstIndex])

	if firstIndex != 0 {
		buf.WriteByte(',')
	}

	s = s[firstIndex:]

	for i := 0; i < len(s);  {
		buf.WriteString(s[i:i + 3])
		i += 3
		if i != len(s) {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	var s string = "123456"
	fmt.Println(comma(s))
}
