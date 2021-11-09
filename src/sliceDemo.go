package main

import "fmt"

var months = [13]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func main() {
	summer := months[6:9]
	fmt.Println(summer)
	fmt.Printf("len:%d, cap:%d\n", len(summer), cap(summer))
	fmt.Printf("val1:%d, val2:%d\n", &months[6], &summer[0])

	fmt.Println("---------------")
	ips := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
	sliceAppend(ips)
	fmt.Println("---------------")
}

func sliceAppend(ips []string) {
	fmt.Printf("len:%d, cap:%d\n", len(ips), cap(ips))
	ips = append(ips, "192.168.0.4", "192.168.0.5", "192.168.0.6")
	ips = append(ips, "192.168.0.7")
	fmt.Printf("len:%d, cap:%d\n", len(ips), cap(ips))
}

