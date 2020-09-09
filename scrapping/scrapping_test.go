package scrapping

import (
	"testing"
)

func TestParseCsv(t *testing.T) {
	_, err := ParseCSV("../url.csv")
	if err != nil {
		t.Errorf("Erreur au moment de parser le fichier : %v\n", err)
	}
}

func TestCreateHTMLFile(t *testing.T) {
	_, err := createHTMLFile("test_filename")
	if err != nil {
		t.Errorf("Erreur sur la création du fichier html : %v\n", err)
	}
}

func TestReadURL(t *testing.T) {
	_, err := readURL("https://www.google.com/")
	if err != nil {
		t.Errorf("Erreur sur la requete GET : %v\n", err)
	}
}

/*
func TestFilterUrlName(t *testing.T) {
	//test debug
	_ = filterURLName("https://www.google.com")
} */
