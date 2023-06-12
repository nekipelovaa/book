package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s = s + sep + arg
		sep = "\n"
	}
	fmt.Println(s)
	fmt.Println(time.Since(start))
	/*-----------------------*/
	start = time.Now()
	fmt.Println(strings.Join(os.Args[:], "\n"))
	fmt.Println(time.Since(start))
}
