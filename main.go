package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"

)

func organizeLinks() (map[string]string) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://feeds.arstechnica.com/arstechnica/index?format=xml")
	if err != nil {
		fmt.Print(err)
	}
	var links string
	var title string

	m := make(map[string]string)

	for _, val := range feed.Items[:5] {
		title = val.Title
		links = val.Link
		m[title] = links
	}
  return m
}

func createCSV() {
  m := organizeLinks()

  csvFile, err := os.Create("email-send/file.csv")

	if err != nil {
		fmt.Print(err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for key, value := range m {
    fmt.Print(key)
    fmt.Print(value)
		r := make([]string, 0, 1+len(value))
		r = append(r, key)
		r = append(r, value)
		err := csvwriter.Write(r)
		if err != nil {
			fmt.Print(err)
		}
	}
  csvwriter.Flush()
}

func main() {
  createCSV()
}
