package main

import "fmt"

// Função de adição
func Adicionar(a, b float64) float64 {
	return a + b
}

// Função de subtração
func Subtrair(a, b float64) float64 {
	return a - b
}

// Função de multiplicação
func Multiplicar(a, b float64) float64 {
	return a * b
}

// Função de divisão
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisão por zero não é permitida")
	}
	return a / b, nil
}

func main() {
	a := 10.0
	b := 5.0

	// Operações básicas
	soma := Adicionar(a, b)
	subtracao := Subtrair(a, b)
	multiplicacao := Multiplicar(a, b)
	divisao, err := Dividir(a, b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Soma: %f\n", soma)
		fmt.Printf("Subtração: %f\n", subtracao)
		fmt.Printf("Multiplicação: %f\n", multiplicacao)
		fmt.Printf("Divisão: %f\n", divisao)
	}
}