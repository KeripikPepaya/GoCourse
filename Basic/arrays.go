package basic

import "fmt"

func main() {

	fmt.Println("Hello")
	var number [5]int
	fmt.Println(number)

	oriArr := [3]int{1, 2, 3}
	copyArr := oriArr

	copyArr[0] = 100
	fmt.Println("original array", oriArr)
	fmt.Println("copy array", copyArr)

	var buah = [3]string{"pisang", "nanas", "pepaya"}
	for index, value := range buah {
		fmt.Println("index", index, "value", value)
	}

	//multidimensi array
	var matrik [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, v := range matrik {
		fmt.Println("value: ", v)
	}

}
