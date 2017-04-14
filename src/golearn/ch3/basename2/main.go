package main

import (
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("basename: %s\n", basename(os.Args[1]))
	}
}
