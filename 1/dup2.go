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
	if indexOfFile != 0 {
		f := false
		for i, s := range files {
			if (1<<i)&indexOfFile != 0 {
				if f {
					fmt.Print(", ", s)
					continue
				}
				fmt.Print(s)
				f = true
			}
		}
	}
}

// Подсчет строк с сохранением в counts
func countLines(f *os.File, counts map[string][2]int, indexOfFile int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// counts[input.Text()][0]++
		// counts[input.Text()][1] |= indexOfFile
		temp := counts[input.Text()]
		temp[0]++
		temp[1] |= indexOfFile
		counts[input.Text()] = temp
		// if counts[input.Text()] == [2]int{0, 0} {
		// 	counts[input.Text()] = [2]int{1, indexOfFile}
		// } else {
		// 	d := counts[input.Text()]
		// 	d[0]++
		// 	d[1] |= indexOfFile
		// counts[input.Text()][0]++
		// counts[input.Text()][1] |= indexOfFile
		//}

	}
}
