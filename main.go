package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error : %v\n", err)
	}
}

func createHTMLFile(filename string, HTMLFileChan chan *os.File) {
	HTMLFile, err := os.Create(filename + ".html")
	checkError(err)
	HTMLFileChan <- HTMLFile
}

func readURL(url string, resBodyChan chan *http.Response) {
	res, err := http.Get(url)
	checkError(err)
	resBodyChan <- res
}

func main() {
	//Debut chrono
	begin := time.Now()

	fileName := flag.String("f", "", "Le nom du fichier qui accueillera notre scrap")
	urlLink := flag.String("url", "", "L'url que l'on va scrapper")
	flag.Parse()

	//On créer nos channel pour gérer nos goroutines
	HTMLFileChan := make(chan *os.File)
	resBodyChan := make(chan *http.Response)

	//goroutines => création fichier html + scrapping de données
	go createHTMLFile(*fileName, HTMLFileChan)
	go readURL(*urlLink, resBodyChan)

	//Récup des données des goroutines
	htmlFile := <-HTMLFileChan
	result := <-resBodyChan
	defer result.Body.Close()

	//On écrit le contenu de notre scraping dans le fichier html
	_, err := io.Copy(htmlFile, result.Body)
	checkError(err)

	//Temps d'éxecution du script
	fmt.Printf("Ce script s'est éxecuté en %v\n", time.Now().Sub(begin))
}
