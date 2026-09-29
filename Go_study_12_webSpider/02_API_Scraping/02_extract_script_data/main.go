package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "view-source:https://quotes.toscrape.com/js/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := http.DefaultClient
	rep, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rep.Body.Close()

	fmt.Println("response Status:", rep.Status)

	doc, err := goquery.NewDocumentFromReader(rep.Body)
}
