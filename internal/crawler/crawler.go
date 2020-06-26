package crawler

import "fmt"

type Crawler int

func (c *Crawler) Respond(seed string, ack *bool) error {
	fmt.Println("Node Crawler is listening and ready to crawl frontiers...")
	return nil
}