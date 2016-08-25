package main

import (
	rss "github.com/jteeuwen/go-pkg-rss"
	"testing"
)

func TestUpdateCache(t *testing.T) {
	fakeItem := &rss.Item{}

	if cacheLen := len(Cache); cacheLen != 0 {
		t.Error("Expected 0 on cache, got ", cacheLen)
	}

	updateCache(12345, fakeItem)
	if cacheLen := len(Cache); cacheLen != 1 {
		t.Error("Expected 1 on cache, got ", cacheLen)
	}
	if indexLen := len(Index); indexLen != 1 {
		t.Error("Expected 1 on cache, got ", indexLen)
	}

	updateCache(12345, fakeItem)
	if cacheLen := len(Cache); cacheLen != 2 {
		t.Error("Expected 2 on cache, got ", cacheLen)
	}
	if indexLen := len(Index); indexLen != 2 {
		t.Error("Expected 2 on cache, got ", indexLen)
	}
	if _, ok := Cache[12346]; !ok {
		t.Error("Expected incremented value on cache")
	}
	if index0, index1 := Index[0], Index[1]; index0 == 12345 && index1 == 12346 {
		t.Error("Expected base and incremented values on index")
	}
}

func TestItemHandler(t *testing.T) {
	fakeItem1 := &rss.Item{PubDate: "Tue, 25 Nov 2014 00:00:00 +0000"}
	fakeItem1pubdate, _ := fakeItem1.ParsedPubDate()
	fakeItem2 := &rss.Item{PubDate: "Tue, 25 Nov 2015 00:00:00 +0000"}
	fakeItem2pubdate, _ := fakeItem2.ParsedPubDate()
	fakeItem3 := &rss.Item{PubDate: "Tue, 25 Nov 2015 00:00:00 +0000"}
	fakeItem3pubdate, _ := fakeItem3.ParsedPubDate()
	fakeSlice := []*rss.Item{fakeItem1, fakeItem2, fakeItem3}
	timeSlice := []int64{fakeItem3pubdate.Unix() + 1, fakeItem2pubdate.Unix(), fakeItem1pubdate.Unix()}

	itemHandler(&rss.Feed{}, &rss.Channel{}, fakeSlice)

	for idx, value := range timeSlice { // reversed
		if _, ok := Cache[value]; !ok {
			t.Error("Expected ", value, " on Cache")
		}

		if v := Index[idx]; v != value {
			t.Error("Expected ", v, " on Index, got", value)
		}
	}
}
