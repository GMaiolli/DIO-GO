package main

import "testing"

func TestSum(t *testing.T) {
	teste := sum(3, 2, 1)
	res := 6

	if teste !=  res {
		t.Error("Valor esperado:", res, "Valor retornado:", teste)
	}
}

func TestSub(t *testing.T) {
	teste := sub(67, 6, 7)
	res := 54

	if teste !=  res {
		t.Error("Valor esperado:", res, "Valor retornado:", teste)
	}
}

func TestMult(t *testing.T) {
	teste := mult(10, 20)
	res := 200

	if teste !=  res {
		t.Error("Valor esperado:", res, "Valor retornado:", teste)
	}
}

func TestDiv(t *testing.T) {
	teste := div(69, 23)
	res := 3

	if teste !=  res {
		t.Error("Valor esperado:", res, "Valor retornado:", teste)
	}
}