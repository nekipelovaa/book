package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][2]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, 0)
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, 1<<i)
			f.Close()
		}

	}
	indexOfFile := 0
	for line, n := range counts {
		if n[0] > 1 {
			fmt.Printf("%d\t%s\n", n[0], line)
			indexOfFile |= n[1]
		}
	}
	for i, s := range files {
		if (1<<i)&indexOfFile != 0 {
			fmt.Print(s)
		}
	}
}
func countLines(f *os.File, counts map[string][2]int, indexOfFile int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// counts[input.Text()][0]++
		// counts[input.Text()][1] |= indexOfFile
		if counts[input.Text()] == [2]int{0, 0} {
			counts[input.Text()] = [2]int{1, indexOfFile}
		} else {
			counts[input.Text()][0]++
			counts[input.Text()][1] |= indexOfFile
		}

	}
}
