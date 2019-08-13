package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func insertion(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, "Bubble sort", bubble([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, "Insertion sort", insertion([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, "Bubble sort", bubble([]int{6, 5, 3, 1, 8, 7, 2, 4}))
	fmt.Fprintln(out, "Insertion sort", insertion([]int{6, 5, 3, 1, 8, 7, 2, 4}))
}
