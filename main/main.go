package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Customer struct {
	Name           string
	FavouriteSnack string
	TotalSnacks    int
}

func main() {
	webPage := "https://candystore.zimpler.net/#candystore-customers"

	res, networkErr := http.Get(webPage)
	if networkErr != nil {
		log.Panic("Failed to connect to server")
	}

	doc, parseErr := goquery.NewDocumentFromReader(res.Body)
	if parseErr != nil {
		log.Panic("Failed to parse data as html")
	}

	var topCustomers []Customer

	// Scrape expected data!
	doc.Find(".summary").First().Find("tbody").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		attrTotalCandy, exists := s.Children().First().Attr("x-total-candy")

		if exists {
			name := s.Children().First().Text()
			totalSnacksInt, _ := strconv.Atoi(attrTotalCandy)
			favSnacks := s.Children().Last().Text()

			c := Customer{
				Name:           name,
				TotalSnacks:    totalSnacksInt,
				FavouriteSnack: favSnacks,
			}

			topCustomers = append(topCustomers, c)
		}
	})

	sortCustomers(topCustomers)
	jsn, err := transformtoJson(topCustomers)

	if err == nil {
		fmt.Println(string(jsn))
	} else {
		log.Panic("Failed to transform to json")
	}
}

func sortCustomers(topCustomers []Customer) {
	sort.Slice(topCustomers, func(i, j int) bool {
		return topCustomers[i].TotalSnacks > topCustomers[j].TotalSnacks
	})
}

func transformtoJson(topCustomers []Customer) ([]byte, error) {
	if jsn, err := json.Marshal(topCustomers); err == nil {
		return jsn, nil
	} else {
		return nil, err
	}
}
