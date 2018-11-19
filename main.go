package main

import (
	"fmt"
)

func main() {
	value := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	var total int
	for i := 0; i < len(value); i++ {
		total += value[i]
	}
	fmt.Println("total value ", total)
	fmt.Println("number of length values ", len(value))
	fmt.Println("average ", total/len(value))
}
