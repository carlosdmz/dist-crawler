package crawler

import (
	"testing"
)

func TestCrawler_Respond(t *testing.T) {
	c := Crawler(0)
	reply := &Reply{Status: "", Data: ""}

	// Get
	c.Respond("https://python.org", reply)
	want := "Node Crawler is listening and ready to crawl its initial seeds like: https://python.org"
	got := reply.Status

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCrawler_Request(t *testing.T) {
	var url = "https://python.org"
	c := Crawler(0)

	reply := &Reply{Status: "", Data: ""}

	c.Respond(url, reply)
	c.Request(url, reply)
	want := "OK"
	got := reply.Status

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCrawler_Crawl(t *testing.T) {
	var url = "https://python.org"
	c := Crawler(0)

	reply := &Reply{Status: "", Data: ""}
	crawlReply := &CrawlReply{"",[]string{}}

	c.Respond(url, reply)
	c.Request(url, reply)
	c.Crawl(reply.Data, crawlReply)
	want := "OK"
	got := reply.Status

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}