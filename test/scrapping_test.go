package test

import (
	"testing"
	"webscrap/scrapping"
)

func TestParseCsv(t *testing.T) {
	_, err := scrapping.ParseCSV("../url.csv")
	if err != nil {
		t.Errorf("Erreur au moment de parser le fichier : %v\n", err)
	}
}
