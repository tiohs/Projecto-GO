package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(1, 2)
	if total != 3 {
		t.Errorf("Resultado esperado da soma Soma(1, 2) = %d, deveria ser 3", total)
	}

}
