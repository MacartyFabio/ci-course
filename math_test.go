package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := Sub(15, 14)

	if total != 1 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}

func TesteMultiplicacao(t *testing.T){
	total := Multiplicacao(2,3)

	if total != 6 {
		t.Errorf("Resultado da Multiplicação é inválido: Resultado %d. Esperado %d, total, 6)	
	}
}
