package main

import (
	"fmt"
	"github.com/cznic/sortutil"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
	"net/http"
	"os"
	"sort"
	"time"
)

var Cache = make(map[int64]*rss.Item)
var Index sortutil.Int64Slice

func handler(w http.ResponseWriter, r *http.Request) {
	var response string

	for _, v := range Index {
		item := Cache[v]
		pubdata, _ := item.ParsedPubDate()
		response += "<b>Title:</b>\n" + item.Title + "\n<b>Data:</b> " + pubdata.String() + "<hr>"
	}

	fmt.Fprintf(w, response)
}

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
		pubdata, _ := item.ParsedPubDate()
		fmt.Println(pubdata.Unix(), pubdata.Unix() + 1)

		Cache[pubdata.Unix()] = item
		Index = append(Index, pubdata.Unix())
	}

	sort.Sort(sort.Reverse(Index))
}

func main() {
	fmt.Println("Serving on: http://localhost:8001/")

	go PollFeed("https://raelmax.github.io/rss.xml", 5, nil)
	go PollFeed("http://hersonls.com.br/rss", 5, nil)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
