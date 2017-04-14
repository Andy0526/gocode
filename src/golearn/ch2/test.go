package main

import "fmt"

var p *int

func main() {
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
	x, y = 1, 1
	fmt.Println(&x == &y, &x == nil)
}
