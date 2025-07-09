package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

func getGoogleTrends() *http.Response {
	url := "https://trends.google.com/trending/rss?geo=AE"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}
	return resp
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}
	return data
}

func main() {
	var rss RSS
	data := readGoogleTrends()

	err := xml.Unmarshal(data, &rss)

	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		os.Exit(1)
	}

	fmt.Println("\n Below are the top trending searches in the UAE for today:\n")
	fmt.Println(("----------------------------------------------------------------"))

	for i := range rss.Channel.ItemList {

		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search term:", rss.Channel.ItemList[i].Title)
		fmt.Println("Link to the trend: ", rss.Channel.ItemList[i].Link)
		fmt.Println("Headline: ", rss.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("Link to the article: ", rss.Channel.ItemList[i].NewsItems[0].HeadlineLink)
		fmt.Println("-----------------------------------------------------------------")
	}
}
