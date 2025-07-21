package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main() {
	type Stock struct {
		company string
		price   string
		change  string
	}
	ticker := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"COST",
		"MCD",
		"V",
		"WMT",
		"DIS",
		"MMM",
		"INTC",
		"AXP",
		"AAPL",
		"BA",
		"CSCO",
		"GS",
		"JPM",
		"CRM",
		"VZ",
	}
	stocks := []Stock{}
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"),
	)
	httpTransport := &http.Transport{
		DisableKeepAlives: true,
	}
	customClient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: httpTransport,
	}
	c.WithTransport(httpTransport)
	c.SetClient(customClient)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*yahoo.*",
		Delay:       2 * time.Second,
		RandomDelay: 1 * time.Second,
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting :", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong", err)
	})
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}
		stock.company = e.ChildText("h1")
		fmt.Println("Company:", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.change)
		stocks = append(stocks, stock)
	})
	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 100},
	)
	for _, t := range ticker {
		q.AddURL("https://finance.yahoo.com/quote/" + t + "/")
	}
	q.Run(c)
	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatal("Failed to create a CSV file to store stock data", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	headers := []string{
		"Company",
		"Price",
		"Change",
	}
	writer.Write(headers)
	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}
