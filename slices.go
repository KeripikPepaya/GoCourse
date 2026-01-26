package main

import "fmt"

func main() {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner_len := i + 1
		twoD[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
