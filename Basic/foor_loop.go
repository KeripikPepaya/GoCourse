package basic

import "fmt"

func main() {
	row := 10

	for i := 1; i <= row; i++ {
		for j := 1; j <= row-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
