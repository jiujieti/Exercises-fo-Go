// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type countAndFiles struct {
	count int
	names []string
}

func main() {
	counts := make(map[string]countAndFiles)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			for _, name := range n.names {
				if name != "" {
					fmt.Println(n.count, name, line)
				} else {
					fmt.Println(n.count, line)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]countAndFiles, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		x := counts[input.Text()]
		x.count++
		x.names = append(x.names, filename)
		counts[input.Text()] = x
	}
}

//!-
