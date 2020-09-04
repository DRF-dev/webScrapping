package main

import (
	"testing"
)

func TestParseCsv(t *testing.T) {
	_, err := ParseCsv("./test.csv")
	if err != nil {
		t.Errorf("Erreur au moment de parser le fichier : %v\n", err)
	}
}
