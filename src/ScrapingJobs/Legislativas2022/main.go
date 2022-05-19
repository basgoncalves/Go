// codigo usado obter resultados das eleicoes por freguesia extraindo os dados do site do MAI
// https://www.legislativas2022.mai.gov.pt/resultados/territorio-nacional?local=LOCAL-010103

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// cria var types que sejam da mesma estrutura do Json file do URL das eleicoes
// usar https://jsonlint.com/ para ver a estrutura Json
type Results struct {
	CurrentResults Current `json:"currentResults"`
	Freguesia      string  `json:"territoryFullName"`
}

type Current struct {
	Parties []PartyResult `json:"resultsParty"`
}

type PartyResult struct {
	Acronym    string  `json:"acronym"`
	Percentage float64 `json:"percentage"`
	Votes      int     `json:"votes"`
}

func main() {

	fmt.Println("A pedir dados ao servidor 'https://www.legislativas2022.mai.gov.pt'...")

	// Registar o tempo inical para calcular o tempo que demorou a correr
	t := time.Now()

	// Output csv filename
	csvFilename := "Resultados_Legislativas_2022.csv"

	file, err := os.Create(csvFilename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	// despeja tudo para o ficheiro
	defer writer.Flush()

	csvFile, err := os.Open("freguesias.csv")
	if err != nil {
		fmt.Println(err)
	}

	areaCodeReader, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	header := []string{"Distrito", "Concelho", "Freguesia",
		"AreaCode", "Party", "Percentage", "Votes"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("A recolher dados por freguesia. Freguesias recolhidas...")
	for i, lineAreaCodeReader := range areaCodeReader {
		if i < 1 {
			continue
		}

		areaCode := lineAreaCodeReader[4]
		// fmt.Println(areaCode)

		if len(areaCode) < 6 {
			areaCode = fmt.Sprintf("0" + areaCode)
		}

		distrito := lineAreaCodeReader[1]
		concelho := lineAreaCodeReader[2]

		if i%100 == 0 {
			fmt.Printf("%d de %d \n", i, len(areaCodeReader))
		}
		results := LoadResults(areaCode) // Grab the results from the URL\
		// fmt.Println(results)
		// time.Sleep(20 * time.Second)
		for _, party := range results.CurrentResults.Parties {
			row := []string{distrito, concelho, results.Freguesia, areaCode, party.Acronym, strconv.FormatFloat(party.Percentage, 'f', 2, 64), strconv.Itoa(party.Votes)}
			err := writer.Write(row)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	fmt.Println("")
	fmt.Println("Dados recolhidos e guardados em ", csvFilename, "Time taken:")
	fmt.Println(time.Since(t))
}

// get data from legislativas2019 url
func LoadResults(areaCode string) Results {
	// URL 2019
	// url := fmt.Sprintf("https://www.eleicoes.mai.gov.pt/legislativas2019/static-data/territory-results/TERRITORY-RESULTS-LOCAL-%s-AR.json", areaCode)

	// URL 2022
	url := fmt.Sprintf("https://www.legislativas2022.mai.gov.pt/frontend/data/TerritoryResults?territoryKey=LOCAL-%s&electionId=AR", areaCode)
	// fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	// Close the request
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var results Results
	json.Unmarshal(body, &results)

	return results
}
