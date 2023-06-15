package main

import "fmt"

func main() {
	c := make(map[string][2]int)
	// if c["s"] == [2]int{0, 0} {
	// d := c["s"]
	// d[0]++

	// c["s"] = d
	c["s"] = [2]int{0, 0}
	c["s"][0] = 1
	// }
	fmt.Println(fmt.Sprint(c["s"][0]))
	//fmt.Println(fmt.Sprint(d))
}
