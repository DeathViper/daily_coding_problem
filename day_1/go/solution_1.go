package main

import (
	"fmt"
)

func quickSort(input []int, start, end int) {
	if len(input) <= 0 || start >= end {
		return
	}
	pivot := input[end]
	j := start - 1
	for i := start; i <= end; i++ {
		if input[i] >= pivot {
			j++
			input[j], input[i] = input[i], input[j]
		}
	}
	quickSort(input, start, j-1)
	quickSort(input, j+1, end)
}

func main() {
	input := []int{10, 15, 3, 7, 2, 5, 7, 9, 45}
	qry := 1700
	quickSort(input, 0, len(input)-1)
	for i, j := 0, len(input)-1; i < j; {
		sum := input[i] + input[j]
		if sum == qry {
			fmt.Println(input[i], input[j])
			return
		} else if sum > qry {
			i++
		} else {
			j--
		}
	}
	fmt.Println(-1)
}
