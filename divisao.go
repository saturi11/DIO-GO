package main

import (
	"fmt"
	"strconv"
)

func main() {
	//i := 1
	//meuSlice := make([]int, 0, 100)
	//for i <= 100 {
	//	if i%3 == 0 {
	//		meuSlice = append(meuSlice, i)
	//	}
	//	i++
	//}
	//fmt.Printf("%v", meuSlice

	i := 1
	meuSlice := make([]string, 0, 100)
	for i <= 100 {
		if i%5 == 0 && i%3 == 0 {
			meuSlice = append(meuSlice, "pin/pan")
		}
		if i%3 == 0 {
			meuSlice = append(meuSlice, "pin")
		} else if i%5 == 0 {
			meuSlice = append(meuSlice, "pan")
		} else {
			a := strconv.Itoa(i)
			meuSlice = append(meuSlice, a)
		}
		i++
	}
	fmt.Printf("%v", meuSlice)

}
