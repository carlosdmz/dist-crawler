package master

import (
	"log"
	"net/rpc"
	"net/url"

	"github.com/carlosdamazio/dist-crawler/pkg/crawler"
	"github.com/carlosdamazio/dist-crawler/pkg/utils"
)

func InitMaster(addr string, seed string) {
	var reply crawler.Reply
	var crawlReply crawler.CrawlReply
	var frontier []string
	var visitedDomains []string
	visitedUrl := seed
	iterations := 50

	log.Println("Master started.")
	node := callNode(addr)

	err := node.Call("Crawler.Respond", visitedUrl, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.Status)

	for iterations > 0 {
		url, err := url.Parse(visitedUrl)
		domain := "http://" + url.Hostname()
		if err != nil {
			log.Fatal(err)
		}

		if utils.FindDuplicateDomain(domain, visitedDomains) {
			visitedUrl, frontier = frontier[0], frontier[1:]
			iterations--
		} else {

			err = node.Call("Crawler.Request", domain, &reply)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(reply.Status)
			visitedDomains = append(visitedDomains, domain)
			log.Println(visitedDomains)

			err = node.Call("Crawler.Crawl", reply.Data, &crawlReply)
			if err != nil {
				log.Fatal(err)
			}
			frontier = append(frontier, crawlReply.Data...)
			visitedUrl, frontier = frontier[0], frontier[1:]
			log.Println(len(frontier))
			iterations--
		}
	}
}

func electNode() {
	log.Println("Electing a node...")
}

func callNode(addr string) *rpc.Client {
	// Might elect a node to seed...
	electNode()

	log.Println("Checking if a node responds...")
	node, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return node
}
