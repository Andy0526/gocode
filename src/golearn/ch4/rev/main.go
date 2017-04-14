package main

import "fmt"

func reverse(arry []int) {
	for i, j := 0, len(arry)-1; i < j; i, j = i+1, j-1 {
		arry[i], arry[j] = arry[j], arry[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:2])
	fmt.Println(a)
	reverse(a[2:])
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

}
