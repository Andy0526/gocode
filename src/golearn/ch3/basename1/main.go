package main

import (
	"fmt"
	"os"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	if len(os.Args) > 1 {
		s_path := os.Args[1]
		s_basename := basename(s_path)
		fmt.Printf("basename: %s\n", s_basename)
	}
}
