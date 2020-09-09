package scrapping

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

//Async : Variable qui nous permet d'attendre nos goroutines
var Async sync.WaitGroup

func createHTMLFile(filename string) (*os.File, error) {

	mkdirName := "html"
	err := os.MkdirAll(mkdirName, os.ModePerm)
	if err != nil {
		return nil, err
	}

	HTMLFile, err := os.Create(mkdirName + "/" + filename + ".html")
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

func filterURLName(urlNotFiltered string) string {
	list := strings.Split(urlNotFiltered, "/")
	domainName := list[2]
	arrayDomainName := strings.Split(domainName, ".")
	if arrayDomainName[0] == "www" {
		return arrayDomainName[1]
	}
	return arrayDomainName[0]
}

//CopyScrapping : fonction copiant dans un fichier HTML les page html scrappé
func CopyScrapping(i int, url string) {
	defer Async.Done()
	name := filterURLName(url)

	htmlFile, err := createHTMLFile(name)
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
}
