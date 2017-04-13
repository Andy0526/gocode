package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	sorted_values := appendValues(values[:0], root)
	return sorted_values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	values := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sorted_values := Sort(values)
	fmt.Printf("%v\n", sorted_values)
}
