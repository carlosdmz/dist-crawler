package crawler

import (
	"testing"
)

func TestCrawler_Respond(t *testing.T) {
	crawler := new(Crawler)
	ack := true
	get := crawler.Respond("", &ack)

	if get != nil {
		t.Error("Respond should not produce any kind of errors whatsoever.")
	}
}