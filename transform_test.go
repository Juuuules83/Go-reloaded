package main

import "testing"

func TestProcessText(t *testing.T) {
	result, err := ProcessText("Hello (up)")

	if err != nil {
		t.Fatal(err)
	}

	if result != "HELLO" {
		t.Errorf("Résultat attendu : HELLO, obtenu : %s", result)
	}
}