package main

import "fmt"

func main() {
	arr := []int{}
	for(i:=0; i < 10; i++) {
		arr := append(arr, i)
	}

	fmt.Println(arr)
}
