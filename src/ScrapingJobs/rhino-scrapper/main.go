// expmple scraping app from https://www.youtube.com/watch?v=4VSno5bK9Uk&t=158s&ab_channel=DivRhino
// uses colly package
// if the "explorer" in VScode is not se to the current directory the "import ("github...")" will be poping up a warning

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"ID"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretreiver.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("could not get an ID")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}
		allFacts = append(allFacts, fact)
		writeJSON(allFacts)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		println(err)
	}

	_ = ioutil.WriteFile("facts.json", file, 0644)

}
