package main

import (
	"testing"
)

func TestAdicionar(t *testing.T) {
	resultado := Adicionar(5, 3)
	if resultado != 8 {
		t.Errorf("A adição de 5 + 3 deveria ser 8, mas obtivemos %f", resultado)
	}
}

func TestSubtrair(t *testing.T) {
	resultado := Subtrair(10, 3)
	if resultado != 7 {
		t.Errorf("A subtração de 10 - 3 deveria ser 7, mas obtivemos %f", resultado)
	}
}

func TestMultiplicar(t *testing.T) {
	resultado := Multiplicar(4, 2)
	if resultado != 8 {
		t.Errorf("A multiplicação de 4 * 2 deveria ser 8, mas obtivemos %f", resultado)
	}
}

func TestDividir(t *testing.T) {
	resultado, err := Dividir(8, 4)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado != 2 {
		t.Errorf("A divisão de 8 / 4 deveria ser 2, mas obtivemos %f", resultado)
	}

	_, err = Dividir(5, 0)
	if err == nil {
		t.Error("A divisão por zero deveria gerar um erro, mas não gerou.")
	}
}
