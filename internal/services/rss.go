package services

import (
	"encoding/xml"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (sctx *ServiceContext) UrlToFeed(url string) (RSSFeed, error) {
	rssFeed := RSSFeed{}
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return rssFeed, err
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	err = decoder.Decode(&rssFeed)
	if err != nil {
		return rssFeed, err
	}
	return rssFeed, nil
}
