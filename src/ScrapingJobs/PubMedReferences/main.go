// this app is designed to import data from all the paper from authors with a "profile" on pub med
// after scrapping the pubmed links, the code creates/updated the "SummaryPapers.csv" with data from the papers of each author
// The code can also be used for single papers (in theory)
//
// Here is an example of people looking for a solution to the problem:
// https://www.researchgate.net/post/Does-anyone-know-a-way-to-populate-Excel-with-user-specified-information-from-a-pubmed-URL-Or-similar-non-excel-method-etc
// https://www.cienciavitae.pt/portal/5411-175F-B4A2

package main

//
import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Author struct {
	Name    string
	Studies []Study
}

type Study struct {
	Title   string
	Year    string
	DOI     string
	PMID    string
	Journal string
}

func main() {

	t := time.Now()

	// Output csv filename
	csvFilename := "SummaryPapers.csv"

	file, err := os.Create(csvFilename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	// despeja tudo para o ficheiro
	defer writer.Flush()

	// read csv file with the names of authors / PUBMED urls
	csvFile, err := os.Open("AuthorList.csv")
	if err != nil {
		fmt.Println(err)
	}

	AuthorList, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	header := []string{"Name", "Title", "Year", "DOI", "PMID", "Journal", "Link"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Collecting data per author...")
	for i, lineAuthorList := range AuthorList {
		if i < 1 {
			continue
		}

		fmt.Println(" ")
		fmt.Println("Loading data for " + lineAuthorList[0])

		urlAuthor := lineAuthorList[1]
		authorData := ScrapeURLS(urlAuthor) // Grab the results from the URL
		authorData.Name = lineAuthorList[0]

		// var authorData Author
		// authorData.Name = lineAuthorList[0]
		// study := Study{}
		// study.Title = "t"
		// study.Year = 2022
		// study.DOI = "doi"
		// study.Journal = "sgtg"
		// study.PMID = "19304"
		// authorData.Studies = append(authorData.Studies, study)

		for _, study := range authorData.Studies {
			row := []string{authorData.Name, study.Title, study.Year, study.DOI, study.Journal}
			err := writer.Write(row)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	fmt.Println("")
	fmt.Println("Data saved in ", csvFilename)
	fmt.Println("Time taken:", time.Since(t))
}

func ScrapeURLS(url string) Author {

	collector := colly.NewCollector()
	authorData := Author{}
	// space := regexp.MustCompile(`\s+`)
	studies := []Study{}
	// for each  URL find the links and data for each title
	link := url
	study := Study{}

	// set up the collector by checking if the link exists / can be opened
	collector.OnHTML("div.docsum-content", func(e *colly.HTMLElement) {

		err := e.Request.Visit(e.Attr("a"))

		t := e.ChildText("a.docsum-title")
		// title := strings.TrimSpace(space.ReplaceAllString(t, " "))
		title := t

		// y := e.ChildText("div.docsum-citation full-citation > span.citation-part")
		year := t
		if strings.Contains(err.Error(), "URL already visited") {
		} else if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println(title)

		study.Title = title
		study.Year = year
		studies = append(studies, study)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})

	// run the function c and whatever is "set-up" inside it
	collector.Visit(link)

	authorData.Studies = studies
	return authorData
}
