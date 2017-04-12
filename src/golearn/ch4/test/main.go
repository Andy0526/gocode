package main

import "fmt"

var c = [...]int{1, 2, 3, 4}

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	b := [...]int{1, 2, 3}

	for i, v := range b {
		fmt.Printf("%d %d\n", i, v)
	}

	r := [...]int{99: -1}
	for i, v := range r {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println(r[0:3])
}

// func main() {
// 	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
// 	Q2 := months[4:7]
// 	summer := months[6:9]
// 	fmt.Println(Q2)
// 	fmt.Println(summer)
// }
