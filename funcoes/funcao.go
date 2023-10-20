package main

import "fmt"

func media(lista []float64) float64 {
	a := lista
	total := 0.0
	for _, valor := range a {
		total += valor
	}
	return total / float64(len(a))
}

func main() {
	lista := []float64{98, 93, 77, 82, 83}
	fmt.Println(media(lista))
}
