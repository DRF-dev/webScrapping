package scrapping

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

//Async : Variable qui nous permet d'attendre nos goroutines
var Async sync.WaitGroup

func createHTMLFile(filename string) (*os.File, error) {
	HTMLFile, err := os.Create(filename + ".html")
	if err != nil {
		return nil, err
	}
	return HTMLFile, nil
}

func readURL(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//CopyScrapping : fonction copiant dans un fichier HTML les page html scrappé
func CopyScrapping(i int, url string) {
	index := strconv.Itoa(i)
	htmlFile, err := createHTMLFile(index)
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier HTML : %v\n", err)
	}

	res, err := readURL(url)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de l'url : %v\n", err)
	}

	defer res.Body.Close()
	//On écrit le contenu de notre scraping dans le fichier html
	_, err = io.Copy(htmlFile, res.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la copy du contenu de l'url : %v\n", err)
	}
	Async.Done()
}
