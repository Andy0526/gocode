package main

import (
	"fmt"
	"os"
	"strconv"

	"go_learn/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, float64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahremheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s=%s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
