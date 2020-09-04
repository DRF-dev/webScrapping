package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//ParseCsv : parse les csv sélectionné
//Prend en paramètre un fichier csv
func ParseCsv(csvFile string) ([]string, error) {
	//On initialise le tableau qui contiendra toutes nos urls
	urls := make([]string, 0)

	//On ouvre le fichier csv et rappel de le fermer à la fin du script
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//On lit chaque url noté dans le csv
	toRead := csv.NewReader(file)
	csvElm, err := toRead.ReadAll()
	if err != nil {
		return nil, err
	}

	//On ajoute ces urls dans le tableau initialisé au début
	for _, elm := range csvElm {
		url := elm[0]
		urls = append(urls, url)
	}

	return urls, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error : %v\n", err)
	}
}

func createHTMLFile(filename string) *os.File {
	HTMLFile, err := os.Create(filename + ".html")
	checkError(err)
	return HTMLFile
}

func readURL(url string) *http.Response {
	res, err := http.Get(url)
	checkError(err)
	return res
}

func main() {
	//Debut chrono
	begin := time.Now()

	fileName := flag.String("f", "", "Le nom du fichier csv")
	flag.Parse()

	//On récupère les html du fichier csv
	urls, err := ParseCsv(*fileName)
	checkError(err)

	//goroutines => création fichier html + scrapping de données
	for i, url := range urls {
		index := strconv.Itoa(i)
		htmlFile := createHTMLFile(index)
		res := readURL(url)
		defer res.Body.Close()
		//On écrit le contenu de notre scraping dans le fichier html
		_, err = io.Copy(htmlFile, res.Body)
		checkError(err)
	}

	//Temps d'éxecution du script
	fmt.Printf("Ce script s'est éxecuté en %v\n", time.Now().Sub(begin))
}
