package main

import "fmt"

var arr [3]int

var q = [3]int{1, 2, 3}

func main() {
	fmt.Println(arr)
	arr[0] = 1
	fmt.Println(arr)

	fmt.Println(q)
	//q = [3]int{2, 3, 4, 5} 不允许，因为长度3和长度4的数组类型是不同的
	fmt.Println(q)

	fmt.Printf("q length is %d\n", len(q))

	paramCopy(q)

	//a := [...]int{1, 2, 3}

	delete()

}

func paramCopy(arr [3]int) {
	fmt.Println("-----------------")
	fmt.Println(arr)
	q = [3]int{3, 4, 5}
	fmt.Println(arr)
	arr = [3]int{7, 8, 9}
	fmt.Println(arr)
	fmt.Println("-----------------")
}
