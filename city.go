package main

import (
	"flag"
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
)


var (
	FeedTitle string
	FeedList []string
	Cache = make(map[int64]*rss.Item)
	Index Int64Slice

	// command line parameters
	port = flag.String("port", "8001", "service port")
	filepath = flag.String("config", "./config.yaml", "config file")
)

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

	t, _ := template.ParseFiles("feed.html")
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

func main() {
	flag.Parse()
	setConfig(*filepath)

	for _, feed := range FeedList {
		go PollFeed(feed, 5, nil)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":" + *port, nil)
}
