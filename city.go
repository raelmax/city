package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/viper"
	"net/http"
	"html/template"
)
var FeedTitle string
var FeedList []string
var Cache = make(map[int64]*rss.Item)
var Index Int64Slice

type Page struct {
	Title string
	Feed []*rss.Item
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


func setConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetDefault("title", "City Feed")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	FeedTitle = viper.GetString("title")
	FeedList = viper.GetStringSlice("feeds")
}

func main() {
	setConfig()

	for _, feed := range FeedList {
		go PollFeed(feed, 5, nil)
	}

	fmt.Println("Serving on: http://localhost:8001/")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
