package main

import "fmt"

func main() {

	arr := [3]int{3, 6, 9}
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
