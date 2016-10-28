package main

import (
	"flag"
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	// "os"
	"strings"
	"strconv"
	"log"
)

var (
	FeedTitle string
	FeedList  []string
	Cache     = make(map[int64]*rss.Item)
	Index     Int64Slice

	port = "8001"
	filepath = "./config.yaml"
	feedtimeout = 5
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

func spawnFeeds() {
	for _, feed := range FeedList {
		go PollFeed(feed, feedtimeout, nil)
	}
}

func parseParams() {
	flag.Parse()

	log.Println(`Usage of ./city:
	-config string
		config file (default "./config.yaml")
	-port string
		service port (default "8001")
	-timeout int
		feed timeout (default 5)
	example: ./city -config=./myconfig.yaml -port=9091 -timeout=15`)

	for _, f := range(flag.Args()) {
		if strings.Index(f, "=") > 0 {
			parts := strings.Split(f, "=")
			switch parts[0] {
			case "-port":
				port = parts[1]
			case "-config":
				filepath = parts[1]
			case "-timeout":
				var err error
				feedtimeout, err = strconv.Atoi(parts[1])
				if err != nil {
					feedtimeout = 5
				}
			}
		}
	}
}

func main() {
	parseParams()

	setConfig(filepath)
	spawnFeeds()

	log.Printf("City starts with params: config=%s, port=%s, timeout=%v\n", filepath, port, feedtimeout)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}