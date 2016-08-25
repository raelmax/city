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
