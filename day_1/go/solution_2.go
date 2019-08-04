package main

import "fmt"

func main() {
	input := []int{10, 15, 3, 7, 2, 5, 7, 9, 45}
	qry := 17
	lookUp := make(map[int]struct{})
	for _, val := range input {
		if _, exists := lookUp[qry-val]; exists {
			fmt.Println(val, qry-val)
			return
		}
		lookUp[val] = struct{}{}
	}
	fmt.Println(-1)
}
