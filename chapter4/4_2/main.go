package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

var num = flag.Int("n", 256, "the num following in 'SHA'")
func main () {
	flag.Parse()
	b := flag.Arg(0)

	switch *num {
	case 256:
		fmt.Println(sha256.Sum256([]byte(b)))
	case 384:
		fmt.Println(sha512.Sum384([]byte(b)))
	case 512:
		fmt.Println(sha512.Sum512([]byte(b)))
	default:
		log.Fatal("error a unexpted num")
	}

}
