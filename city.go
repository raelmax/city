package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"net/http"
)

var Cache = make(map[int64]*rss.Item)
var Index Int64Slice

func handler(w http.ResponseWriter, r *http.Request) {
	var response string

	for _, v := range Index {
		item := Cache[v]
		response += "<b>Title:</b>\n" + item.Title + "\n<b>Description:</b>\n" + item.Description + "<hr>"
	}

	fmt.Fprintf(w, response)
}

func main() {
	fmt.Println("Serving on: http://localhost:8001/")

	go PollFeed("https://raelmax.github.io/rss.xml", 5, nil)
	go PollFeed("http://hersonls.com.br/rss", 5, nil)
	go PollFeed("http://everson.com.br/rss.xml", 5, nil)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
