package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"webscrap/scrapping"
)

func main() {
	//Debut chrono
	begin := time.Now()

	//Création du flag
	fileName := flag.String("f", "", "Le nom du fichier csv")
	flag.Parse()

	//On récupère les html du fichier csv
	urls, err := scrapping.ParseCSV(*fileName)
	if err != nil {
		log.Fatalf("Erreur lors du de la récupération des urls : %v\n", err)
	}

	//goroutines: création fichier html + scrapping de données
	for i, url := range urls {
		scrapping.Async.Add(1)
		go scrapping.CopyScrapping(i, url)
	}
	scrapping.Async.Wait()

	//Temps d'éxecution du script
	fmt.Printf("Ce script s'est éxecuté en %v\n", time.Now().Sub(begin))
}
