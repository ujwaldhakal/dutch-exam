package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"sync"
)

func loadLessonsMap(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Cannot read lessons")
	}
	return data
}

func crawlDataToFile(fileName string, websiteUrl string) {
	data := make(map[string]string)
	fName := "service/lessons/" + fileName + ".json"
	file, err := os.Create(fName)
	if err != nil {

		return
	}
	defer file.Close()
	c := colly.NewCollector()

	c.OnHTML("html", func(e *colly.HTMLElement) {
		goQuery := e.DOM
		firstTable := goQuery.Find("table").First()
		firstTable.Find("tr").Each(func(i int, selection *goquery.Selection) {
			var key, value string
			selection.Find("td").Each(func(index int, td *goquery.Selection) {
				if index == 0 {
					value = td.Text()
				}
				if index == 1 {
					key = td.Text()
				}
			})
			data[key] = value
		})

	})

	c.Visit(websiteUrl)
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func main() {

	jsonSiteMap := loadLessonsMap("service/lessons/lessons-map.json")

	mappedSitemap := make(map[string]string)
	jsonError := json.Unmarshal(jsonSiteMap, &mappedSitemap)
	if jsonError != nil {
		log.Fatal("cannot unmarshall json data")
	}

	var wg sync.WaitGroup
	wg.Add(len(mappedSitemap))
	for fileName, websiteUrl := range mappedSitemap {
		go func(fileName string, websiteUrl string) {
			defer wg.Done()
			crawlDataToFile(fileName, websiteUrl)
		}(fileName, websiteUrl)

	}
	wg.Wait()
}
