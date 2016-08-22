package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
	"os"
	"sort"
	"time"
)

func PollFeed(uri string, timeout int, cr xmlx.CharsetFunc) {
	feed := rss.New(timeout, true, nil, itemHandler)

	for {
		if err := feed.Fetch(uri, cr); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s\n", uri, err)
			return
		}

		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for _, item := range newitems {
		pubdate, _ := item.ParsedPubDate()
		updateCache(pubdate.Unix(), item)
	}
}

func updateCache(key int64, item *rss.Item) {
	if _, ok := Cache[key]; ok {
		updateCache(key + 1, item)
	} else {
		Cache[key] = item
		Index = append(Index, key)
		sort.Sort(sort.Reverse(Index))
	}
}