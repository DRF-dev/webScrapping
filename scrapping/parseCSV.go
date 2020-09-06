package scrapping

import (
	"encoding/csv"
	"os"
)

//ParseCSV : parse les csv sélectionné
func ParseCSV(csvFile string) ([]string, error) {
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
