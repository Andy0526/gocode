package main

import "fmt"

// func appendIdx(x []int, y int) []int {
// 	var z []int
// 	zlen := len(x) + 1
// 	if zlen <= cap(x) {
// 		z = x[:zlen]
// 	} else {
// 		zcap := zlen
// 		if zcap < 2*len(x) {
// 			zcap = 2 * len(x)
// 		}
// 		z = make([]int, zlen, zcap)
// 		copy(z, x)
// 	}
// 	z[len(x)] = y
// 	return z
// }

// func appendIdx(x []int, y ...int) []int {
// 	var z []int
// 	zlen := len(x) + len(y)
// 	if zlen <= cap(x) {
// 		z = x[:zlen]
// 	} else {
// 		zcap := zlen
// 		if zcap < 2*len(x) {
// 			zcap = 2 * len(x)
// 		}
// 		z = make([]int, zlen, zcap)
// 		copy(z, x)
// 	}
// 	copy(z[len(x):], y)
// 	return z
// }

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendIdx(x, []int{11, 12})
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
