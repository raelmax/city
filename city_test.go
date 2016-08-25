package main

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	if feedLen := len(FeedList); feedLen != 0 {
		t.Errorf("Expected 0 to feedLen, got %d", feedLen)
	}
	if FeedTitle != "" {
		t.Errorf("Expected blank string to feedTitle, got %d", FeedTitle)
	}

	setConfig("testdata/test_config.yaml")

	if feedLen := len(FeedList); feedLen != 1 {
		t.Errorf("Expected 1 to feedLen, got %d", feedLen)
	}
	if feedItem := FeedList[0]; feedItem != "https://raelmax.github.io/rss.xml" {
		t.Errorf("Wrong value to feed item from test config")
	}
	if FeedTitle != "Test Config Title" {
		t.Errorf("Expected title from config yaml")
	}
}
