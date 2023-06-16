package services

import (
	"context"
	"database/sql"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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
	// sctx.Logger.Printf("%+v", rssFeed)
	if err != nil {
		sctx.Logger.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubdate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			sctx.Logger.Printf("Couldn't parse date %v with err %v", item.PubDate, err)
			continue
		}

		post := Post{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubdate,
			Url:         item.Link,
			FeedId:      feed.ID,
		}

		err = sctx.CreatePost(context.Background(), post)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			sctx.Logger.Println("Could not create post:", err)
		}
	}
	sctx.Logger.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
