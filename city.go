package main

import (
	"flag"
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	"log"
)

var (
	FeedTitle string
	FeedList  []string
	Cache     = make(map[int64]*rss.Item)
	Index     Int64Slice

	// command line parameters
	port        = flag.String("port", "8001", "service port")
	filepath    = flag.String("config", "./config.yaml", "config file")
	feedtimeout = flag.Int("timeout", 5, "feed timeout")
)

type MyFeedItem struct {
	Title		string
	PubDate		string
	Description	string
	Guid		string
	Categories	[]string
	Links		[]string
}

type Page struct {
	Title string
	Feed  []*MyFeedItem
}

//type Page struct {
//	Title string
//	Feed  []*rss.Item
//}

func handler(w http.ResponseWriter, r *http.Request) {
	feed := []*rss.Item{}

	for _, v := range Index {
		feed = append(feed, Cache[v])
	}

	var newFeed []*MyFeedItem
	for _, f := range(feed) {
		date, err := f.ParsedPubDate()
		parsedDate := ""
		if err == nil {
			parsedDate = date.Format("2006-02-01")
		}

		var categories []string
		for _, cat := range(f.Categories) {
			if cat.Text != "" {
				categories = append(categories, cat.Text)
			}
		}

		var links []string
		for _, link := range(f.Links) {
			links = append(links, link.Href)
		}


		item := MyFeedItem{
			Title: f.Title,
			PubDate: parsedDate,
			Description: f.Description,
			Guid: *f.Guid,
			Categories: categories,
			Links: links,
		}
		newFeed = append(newFeed, &item)

		log.Println(item.Title)
	}

	page := Page{Title: FeedTitle, Feed: newFeed}
	funcMap := template.FuncMap{
		"html": func(value interface{}) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
	}

	t, err := template.ParseFiles("feed.html")
	if err != nil {
		log.Panic(err)
	}
	t.Funcs(funcMap).Execute(w, page)
}

func setConfig(filepath string) {
	viper.SetConfigFile(filepath)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	FeedTitle = viper.GetString("title")
	FeedList = viper.GetStringSlice("feeds")
}

func spawnFeeds() {
	for _, feed := range FeedList {
		go PollFeed(feed, *
			feedtimeout, nil)
	}
}

func main() {
	flag.Parse()
	setConfig(*filepath)
	spawnFeeds()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+*port, nil)
}
