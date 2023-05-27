package services

import (
	"context"
	"sync"
	"time"
)

func (sctx *ServiceContext) StartScraping(concurrency int, timeBetweenRequest time.Duration) {
	sctx.Logger.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := sctx.GetNextFeedsToFetch(context.Background(), concurrency)
		if err != nil {
			sctx.Logger.Println("error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go sctx.scrapeFeed(wg, feed)
		}
		wg.Wait()
	}
}

func (sctx *ServiceContext) scrapeFeed(wg *sync.WaitGroup, feed Feed) {
	defer wg.Done()

	_, err := sctx.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		sctx.Logger.Println("Error markin feed as fetched:", err)
		return
	}

	rssFeed, err := sctx.UrlToFeed(feed.URL)
	if err != nil {
		sctx.Logger.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		sctx.Logger.Println("Found post", item.Title)
	}
	sctx.Logger.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
