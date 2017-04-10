package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			file, ok := counts[line]
			if !ok {
				file = make(map[string]int)
			}
			file[filename]++
		}
	}
	for line, data := range counts {
		if len(data) > 1 {
			for file, n := range data {
				if n > 0 {
					fmt.Printf("%d\t%s\t%s\n", n, file, line)
				}
			}
		}
	}
}
