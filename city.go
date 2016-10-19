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
	layout    = flag.String("layout", "feed", "layout template (HTML file without extension)")
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
	Feed  []*rss.Item
}

func handler(w http.ResponseWriter, r *http.Request) {
	feed := []*rss.Item{}

	for _, v := range Index {
		feed = append(feed, Cache[v])
	}

	page := Page{Title: FeedTitle, Feed: feed}
	funcMap := template.FuncMap{
		"html": func(value interface{}) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
	}

	t, err := template.ParseFiles(*layout + ".html")
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
